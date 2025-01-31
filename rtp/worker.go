package sgc7rtp

import (
	"context"
	"math"
	"time"

	goutils "github.com/zhs007/goutils"
	sgc7game "github.com/zhs007/slotsgamecore7/game"
	"go.uber.org/zap"
	"gonum.org/v1/gonum/stat"
)

// FuncOnRTPTimer - on timer for rtp
type FuncOnRTPTimer func(totalnums int64, curnums int64, curtime time.Duration)

// StartRTP - start RTP
func StartRTP(game sgc7game.IGame, rtp *RTP, worknums int, spinnums int64, stake *sgc7game.Stake, numsTimer int,
	ontimer FuncOnRTPTimer, needVariance bool, limitPayout int64) time.Duration {

	t1 := time.Now()

	lastnums := worknums
	ch := make(chan *RTP)
	chTimer := make(chan int)

	go func() {
		lastspinnums := spinnums

		for {
			curnums := <-chTimer

			lastspinnums -= int64(curnums)

			ontimer(spinnums, spinnums-lastspinnums, time.Since(t1))

			if lastspinnums <= 0 {
				break
			}
		}
	}()

	for i := 0; i < worknums; i++ {
		go func() {
			currtp := rtp.Clone()

			plugin := game.NewPlugin()
			defer game.FreePlugin(plugin)

			ps := game.NewPlayerState()
			// ps := sgc7game.NewBasicPlayerState("bg")
			results := []*sgc7game.PlayResult{}
			cmd := "SPIN"
			off := 0

			for i := int64(0); i < spinnums/int64(worknums); i++ {
				pbsjson := ps.GetPublicJson()
				ppsjson := ps.GetPrivateJson()
				iserrturn := false

				plugin.ClearUsedRngs()

				// currtp.Bet(stake.CashBet)

				totalReturn := int64(0)
				for {
					pr, err := game.Play(plugin, cmd, "", ps, stake, results)
					if err != nil {
						iserrturn = true

						goutils.Error("StartRTP.Play",
							zap.Int("results", len(results)),
							zap.Error(err))

						break
					}

					if pr == nil {
						break
					}

					results = append(results, pr)
					if pr.IsFinish {

						if currtp.Stats2 != nil {
							currtp.Stats2.OnResults(stake, results)
						}

						break
					}

					if pr.IsWait {
						break
					}

					if len(pr.NextCmds) > 0 {
						cmd = pr.NextCmds[0]
					} else {
						cmd = "SPIN"
					}
				}

				if iserrturn {
					ps.SetPublicJson(pbsjson)
					ps.SetPrivateJson(ppsjson)

					results = results[:0]

					i--
				} else {
					if limitPayout > 0 {
						cp := int64(0)
						for _, v := range results {
							if cp+v.CashWin < limitPayout {
								cp += v.CashWin
							} else {
								v.CashWin = limitPayout - cp
								cp = limitPayout
							}
						}
					}

					currtp.Bet(stake.CashBet)

					for _, v := range results {
						currtp.TotalWins += v.CashWin
						totalReturn += v.CashWin

						currtp.OnResult(stake, v)
					}

					currtp.OnResults(results)

					if needVariance {
						currtp.AddReturns(float64(totalReturn / stake.CashBet))
					}

					results = results[:0]

					off++
					if off >= numsTimer {
						chTimer <- off

						off = 0
					}
				}
			}

			currtp.OnPlayerPoolData(ps)

			ch <- currtp
		}()
	}

	for {
		currtp := <-ch

		rtp.Add(currtp)

		lastnums--

		if lastnums <= 0 {
			break
		}
	}

	elapsed := time.Since(t1)

	if needVariance {
		rtp.Variance = stat.Variance(rtp.Returns, rtp.ReturnWeights)
		rtp.StdDev = stat.StdDev(rtp.Returns, rtp.ReturnWeights)
	}

	return elapsed
}

// StartRTP - start RTP
func StartScaleRTPDown(game sgc7game.IGame, rtp *RTP, worknums int, spinnums int64, stake *sgc7game.Stake, numsTimer int,
	ontimer FuncOnRTPTimer, hitFrequency float64, originalRTP float64, targetRTP float64, needVariance bool) time.Duration {

	val := int((targetRTP/originalRTP - hitFrequency) / (1 - hitFrequency) * math.MaxInt32)

	t1 := time.Now()

	lastnums := worknums
	ch := make(chan *RTP)
	chTimer := make(chan int)

	go func() {
		lastspinnums := spinnums

		for {
			curnums := <-chTimer

			lastspinnums -= int64(curnums)

			ontimer(spinnums, spinnums-lastspinnums, time.Since(t1))

			if lastspinnums <= 0 {
				break
			}
		}
	}()

	for i := 0; i < worknums; i++ {
		go func() {
			currtp := rtp.Clone()

			plugin := game.NewPlugin()
			defer game.FreePlugin(plugin)

			ps := game.NewPlayerState()
			// ps := sgc7game.NewBasicPlayerState("bg")
			results := []*sgc7game.PlayResult{}
			cmd := "SPIN"
			off := 0

			for i := int64(0); i < spinnums/int64(worknums); {
				pbsjson := ps.GetPublicJson()
				ppsjson := ps.GetPrivateJson()
				iserrturn := false

				plugin.ClearUsedRngs()

				totalReturn := int64(0)
				for {
					pr, err := game.Play(plugin, cmd, "", ps, stake, results)
					if err != nil {
						iserrturn = true

						goutils.Error("StartScaleRTPDown.Play",
							zap.Int("results", len(results)),
							zap.Error(err))

						break
					}

					if pr == nil {
						break
					}

					results = append(results, pr)
					if pr.IsFinish {
						break
					}

					if pr.IsWait {
						break
					}

					if len(pr.NextCmds) > 0 {
						cmd = pr.NextCmds[0]
					} else {
						cmd = "SPIN"
					}
				}

				if iserrturn {
					ps.SetPublicJson(pbsjson)
					ps.SetPrivateJson(ppsjson)

					results = results[:0]
				} else {
					iswin := false
					for _, v := range results {
						if v.CoinWin > 0 {
							iswin = true

							break
						}
					}

					if iswin {
						cr, err := plugin.Random(context.Background(), math.MaxInt32)
						if err != nil {
							goutils.Error("StartScaleRTPDown.Random",
								zap.Int("results", len(results)),
								zap.Error(err))
						}

						if cr > val {
							results = results[:0]

							continue
						}
					}

					currtp.Bet(stake.CashBet)

					for _, v := range results {
						currtp.TotalWins += v.CashWin
						totalReturn += v.CashWin

						currtp.OnResult(stake, v)
					}

					currtp.OnResults(results)

					if needVariance {
						rtp.AddReturns(float64(totalReturn / stake.CashBet))
						// rtp.Returns = append(rtp.Returns, float64(totalReturn/stake.CashBet))
					}

					results = results[:0]

					off++
					if off >= numsTimer {
						chTimer <- off

						off = 0
					}

					i++
				}
			}

			currtp.OnPlayerPoolData(ps)

			ch <- currtp
		}()
	}

	for {
		currtp := <-ch

		rtp.Add(currtp)

		lastnums--

		if lastnums <= 0 {
			break
		}
	}

	elapsed := time.Since(t1)

	if needVariance {
		rtp.Variance = stat.Variance(rtp.Returns, rtp.ReturnWeights)
	}

	return elapsed
}
