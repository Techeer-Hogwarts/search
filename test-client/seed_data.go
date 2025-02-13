package main

import (
	"fmt"
	"time"

	"github.com/Techeer-Hogwarts/search/config"
	"github.com/meilisearch/meilisearch-go"
)

func main() {
	// Wait for Meilisearch to be ready
	time.Sleep(2 * time.Second)
	// config.LoadEnvFile("../.env")

	meiliHost := config.GetEnvVarAsString("MEILISEARCH_ADDR", "http://localhost:7700")
	meiliKey := config.GetEnvVarAsString("MEILISEARCH_API_KEY", "masterKey")

	client := meilisearch.New(meiliHost, meilisearch.WithAPIKey(meiliKey))

	// Define index settings
	blogIndex := client.Index("blog")
	userIndex := client.Index("user")
	resumeIndex := client.Index("resume")
	sessionIndex := client.Index("session")
	projectTeamIndex := client.Index("projectTeam")
	studyTeamIndex := client.Index("studyTeam")
	eventIndex := client.Index("event")

	// Sample data
	blogDocuments := []map[string]interface{}{
		{"id": "1", "title": "Next.js vs Nest.js", "body": "네스트와 넥스트를 알아보자"},
		{"id": "2", "title": "왜 파이썬은 느린가", "body": "파이썬이 왜 느린지 알아보자"},
	}

	userDocuments := []map[string]interface{}{
		{"id": "1", "user": "윤정은", "title": "사용자 윤정은", "body": "johndoe@example.com"},
		{"id": "2", "user": "김미영", "title": "사용자 김미영", "body": "janesmith@example.com"},
	}

	resumeDocuments := []map[string]interface{}{
		{"id": "1", "title": "윤정은", "body": "윤정은의 이력서"},
		{"id": "2", "title": "김미영", "body": "김미영의 이력서"},
		{"id": "3", "title": "윤정은", "body": "윤정은의 백앤드 이력서"},
	}

	sessionDocuments := []map[string]interface{}{
		{"id": "1", "user": "윤정은", "title": "백앤드 기초", "description": "장고, restAPI, crud"},
		{"id": "2", "user": "김미영", "title": "백앤드 심화 과정", "description": "장고, 스프링, nestjs"},
		{"id": "3", "user": "윤정은", "title": "클라우드와 인프라", "description": "AWS, GCP, Azure"},
	}

	projectTeamDocuments := []map[string]interface{}{
		{"id": "1", "title": "호그와트", "body": "테커의 테커를에 위한 테커를 위한"},
		{"id": "2", "title": "래플", "body": "이커머스, 우린 이렇게 봅니다"},
		{"id": "3", "title": "테커 로그", "body": "테커 프로젝트의 모든 로그를 기록합니다"},
	}

	studyTeamDocuments := []map[string]interface{}{
		{"id": "1", "title": "네스트 스터디", "body": "Nest.js 스터디 팀입니다."},
		{"id": "2", "title": "넥스트 스터디", "body": "Next.js 스터디 팀입니다."},
		{"id": "3", "title": "고랭 스터디", "body": "Go 언어 스터디 팀입니다."},
	}

	eventDocuments := []map[string]interface{}{
		{"id": "1", "title": "테커 파티", "body": "테커 파티는 테커를 위한 파티입니다."},
		{"id": "2", "title": "토스 개발자 세션", "body": "토스 개발자 김재연님의 세션입니다."},
		{"id": "3", "title": "실리콘 밸리 개발자의 모든것", "body": "실리콘 밸리 개발자의 모든것을 알아봅시다."},
	}

	// Insert data
	_, err := blogIndex.AddDocuments(blogDocuments)
	if err != nil {
		fmt.Println("Error inserting blog data:", err)
		return
	}

	_, err = userIndex.AddDocuments(userDocuments)
	if err != nil {
		fmt.Println("Error inserting user data:", err)
		return
	}

	_, err = resumeIndex.AddDocuments(resumeDocuments)
	if err != nil {
		fmt.Println("Error inserting resume data:", err)
		return
	}

	_, err = sessionIndex.AddDocuments(sessionDocuments)
	if err != nil {
		fmt.Println("Error inserting session data:", err)
		return
	}

	_, err = projectTeamIndex.AddDocuments(projectTeamDocuments)
	if err != nil {
		fmt.Println("Error inserting projectTeam data:", err)
		return
	}

	_, err = studyTeamIndex.AddDocuments(studyTeamDocuments)
	if err != nil {
		fmt.Println("Error inserting studyTeam data:", err)
		return
	}

	_, err = eventIndex.AddDocuments(eventDocuments)
	if err != nil {
		fmt.Println("Error inserting event data:", err)
		return
	}

	fmt.Println("Sample data successfully inserted into Meilisearch!")
}
