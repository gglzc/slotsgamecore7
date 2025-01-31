package dtserv

import (
	"context"
	"io"

	goutils "github.com/zhs007/goutils"
	sgc7pbutils "github.com/zhs007/slotsgamecore7/pbutils"
	sgc7pb "github.com/zhs007/slotsgamecore7/sgc7pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Client - DTGameLogicClient
type Client struct {
	servAddr string
	conn     *grpc.ClientConn
	client   sgc7pb.DTGameLogicClient
}

// NewClient - new DTGameLogicClient
func NewClient(servAddr string) (*Client, error) {
	client := &Client{
		servAddr: servAddr,
	}

	return client, nil
}

// reset - reset
func (client *Client) reset() {
	if client.conn != nil {
		client.conn.Close()
	}

	client.conn = nil
	client.client = nil
}

// GetConfig - get config
func (client *Client) GetConfig(ctx context.Context) (*sgc7pb.GameConfig, error) {
	if client.conn == nil || client.client == nil {
		conn, err := grpc.Dial(client.servAddr, grpc.WithInsecure())
		if err != nil {
			goutils.Error("Client.GetConfig:grpc.Dial",
				zap.String("server address", client.servAddr),
				zap.Error(err))

			return nil, err
		}

		client.conn = conn
		client.client = sgc7pb.NewDTGameLogicClient(conn)
	}

	res, err := client.client.GetConfig(ctx, &sgc7pb.RequestConfig{})
	if err != nil {
		goutils.Error("Client.GetConfig:GetConfig",
			zap.String("server address", client.servAddr),
			zap.Error(err))

		client.reset()

		return nil, err
	}

	return res, nil
}

// Initialize - initialize a player
func (client *Client) Initialize(ctx context.Context) (*sgc7pb.PlayerState, error) {
	if client.conn == nil || client.client == nil {
		conn, err := grpc.Dial(client.servAddr, grpc.WithInsecure())
		if err != nil {
			goutils.Error("Client.Initialize:grpc.Dial",
				zap.String("server address", client.servAddr),
				zap.Error(err))

			return nil, err
		}

		client.conn = conn
		client.client = sgc7pb.NewDTGameLogicClient(conn)
	}

	res, err := client.client.Initialize(ctx, &sgc7pb.RequestInitialize{})
	if err != nil {
		goutils.Error("Client.Initialize:Initialize",
			zap.String("server address", client.servAddr),
			zap.Error(err))

		client.reset()

		return nil, err
	}

	return res, nil
}

// Play - play game
func (client *Client) Play(ctx context.Context, ps *sgc7pb.PlayerState,
	cheat string,
	stake *sgc7pb.Stake,
	clientParams string,
	cmd string) (*sgc7pb.ReplyPlay, error) {

	if client.conn == nil || client.client == nil {
		conn, err := grpc.Dial(client.servAddr, grpc.WithInsecure())
		if err != nil {
			goutils.Error("Client.Play:grpc.Dial",
				zap.String("server address", client.servAddr),
				zap.Error(err))

			return nil, err
		}

		client.conn = conn
		client.client = sgc7pb.NewDTGameLogicClient(conn)
	}

	stream, err := client.client.Play(ctx, &sgc7pb.RequestPlay{
		PlayerState:  ps,
		Cheat:        cheat,
		Stake:        stake,
		ClientParams: clientParams,
		Command:      cmd,
	})
	if err != nil {
		goutils.Error("Client.Play:Play",
			zap.String("server address", client.servAddr),
			zap.Error(err))

		client.reset()

		return nil, err
	}

	reply := &sgc7pb.ReplyPlay{}

	for {
		rp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return reply, nil
			}

			goutils.Error("Client.Play:Recv",
				zap.String("server address", client.servAddr),
				zap.Error(err))

			client.reset()

			return nil, err
		}

		if rp != nil {
			sgc7pbutils.MergeReplyPlay(reply, rp)
		}
	}
}
