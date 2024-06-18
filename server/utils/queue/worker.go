package queue

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func Worker() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: global.GVA_CONFIG.Redis.Addr, Password: global.GVA_CONFIG.Redis.Password, DB: global.GVA_CONFIG.Redis.DB},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeEmailDelivery, HandleEmailDeliveryTask)
	mux.Handle(TypeImageResize, NewImageProcessor())
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		global.GVA_LOG.Error("worker start failed!", zap.Error(err))
	}
}
