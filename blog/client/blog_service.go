package main

import (
	"context"
	pb "cristianrb/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Cristian",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---readBlog was invoked---")

	req := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---updateBlog was invoked---")

	blog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Cristian",
		Title:    "New titlte",
		Content:  "Updated content",
	}

	_, err := c.UpdateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Println("Blog has been updated")
}
