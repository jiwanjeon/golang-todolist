package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jiwanjeon/go-todolist/pkg/routes"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTodoListRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}

// TODO : Complete 와 InComplete가 하는 동작이 똑같은데 하나로 합칠 수 있을듯?
// TODO: 아 아니다 프론트 쪽에서 complete 인지 InComplete 인지 상태를 개별적으로 쏴주기 때문에 api관리를 따로 하는게 맞는듯?
// 아니지 그게 바꾸는게 Update니까 말 그대로 업데이트에 넣어주면 되지 않을까?

// TODO with 종민님

/*
	> C.R.U.D Interface 
		
		- Method 명 동사로
		- 모델 객체를 목적어가 아닌 주어로 생각하여 함수 작성 (CreateTodo --> todo *Todo)
		- Migration의 경우 수동으로 하는게 좋음(데이터 손실 방지)
	
	> 변수 이름
		- 객체 이름은 명사형
*/