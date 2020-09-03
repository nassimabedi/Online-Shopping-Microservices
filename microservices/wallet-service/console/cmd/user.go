package cmd

import (
	"Online-Shopping-Microservices/microservices/wallet-service/constant"
	"Online-Shopping-Microservices/microservices/wallet-service/logic"
	"Online-Shopping-Microservices/microservices/wallet-service/repository"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// GetUser represents the base command when called without any subcommands
var GetUser = &cobra.Command{
	Use:   "GetUser",
	Short: "A worker for get users and add to wallet",
	Long:  `A worker for get users from redis queue and add to wallet`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()
		RepoQueue := repository.NewQueueRepo(ctx)
		LogicCreatCredit := logic.NewCreateWallet(ctx)

		i := 0
		winLen := int(constant.WinningUserEnd)
		for {
			//pop user phone number from queue user (discount service push phone number)
			user, err := RepoQueue.PopUserFromQueue()
			if err != nil {
				fmt.Printf("error on pop new data from user queue with error :%v \n", err)
				continue
			}
			//insert userinfo into Permanent database
			_, err = LogicCreatCredit.CreateNewCredit(user, constant.CreditAmount)

			if err != nil {
				fmt.Printf("error on insert credit logic wit error :%v \n", err)
				continue
			}

			time.Sleep(2 * time.Second)

			if i < winLen {
				//push user info in new queue for the purpose of real time user
				_, err = RepoQueue.InsertForRealTime(user)
				if err != nil {
					fmt.Printf("error on push new data from user real time queue with error :%v \n", err)
					continue
				}
			}
			i += 1

			if i > winLen {

				break
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(GetUser)
}
