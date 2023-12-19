package main

import (
	"context"
	"log"
	pb "mailinglist/proto"
	"time"

	"github.com/alexflint/go-arg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func logResponse(res *pb.EmailResponse, err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if res.EmailEntry == nil {
		log.Printf(" email not found")
	} else {
		log.Printf("response: %v", res.EmailEntry)
	}
}

func createEmail(client pb.MailingListServiceClient, addr string) {
	log.Println("create email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.CreateEmail(ctx, &pb.CreateEmailRequest{EmailAddr: addr})
	if err != nil {
		log.Printf("error: %v", err)
	}
	logResponse(res, err)

}

func updateEmail(client pb.MailingListServiceClient, entry pb.EmailEntry) {
	log.Println("update email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.UpdateEmail(ctx, &pb.UpdateEmailRequest{EmailEntry: &entry})
	if err != nil {
		log.Printf("error: %v", err)
	}
	logResponse(res, err)

}

func getEmail(client pb.MailingListServiceClient, addr string) {
	log.Println("get email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.GetEmail(ctx, &pb.GetEmailRequest{EmailAddr: addr})
	if err != nil {
		log.Printf("error: %v", err)
	}
	logResponse(res, err)

}

func getEmailBatch(client pb.MailingListServiceClient, count, page int32) {
	log.Println("get email batch")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.GetEmailBatch(ctx, &pb.GetEmailBatchRequest{Count: count, Page: page})
	if err != nil {
		log.Printf("error: %v", err)
	}

	for _, emailEntry := range res.EmailEntries {
		log.Printf(" EmailEntry: %v\n", emailEntry)
	}

}

func deleteEmail(client pb.MailingListServiceClient, addr string) {
	log.Println("delete email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.DeleteEmail(ctx, &pb.DeleteEmailRequest{EmailAddr: addr})
	if err != nil {
		log.Printf("error: %v", err)
	}
	logResponse(res, err)

}

var args struct {
	GrpcAddr string `arg:"env:MAILINGLIST_GRPC_ADDR"`
}

func main() {
	arg.MustParse(&args)
	if args.GrpcAddr == "" {
		args.GrpcAddr = ":8081"
	}
	conn, err := grpc.Dial(args.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMailingListServiceClient(conn)

	getEmailBatch(client, 10, 10)
	createEmail(client, "999@999.com")

}
