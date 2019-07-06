package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

type DecoratorRetryApiResult struct {
	ApiData  string
	ApiError error
}

// 模拟接口请求发起
func doErrorApiRequest() (string, error) {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	numRand := seededRand.Intn(2)
	// 随机返回异常
	switch numRand {
	case 0:
		// 错误
		return "", errors.New("api error")
	case 1:
		// 超时
		time.Sleep(10 * time.Second)
		return "", errors.New("api timeout")
	default:
		// 正常
		return "1", nil
	}
}

// 接口重试装饰器
func retryApiDecorator(apiFunc func() (string, error), decoratorResultChan chan DecoratorRetryApiResult, countRetry, timeout int) {

	done := make(chan bool)  // 接口请求完成
	retry := make(chan bool) // 接口请求重试

	for i := 0; i < countRetry+1; i++ {

		fmt.Println(i)
		fmt.Println(countRetry)

		// 显示重试次数
		if i > 0 {
			fmt.Println()
			fmt.Println("[request ] request retry:", i)
		}

		// 设置单次请求超时
		d := time.Duration(timeout) * time.Second
		timeAfter := time.After(d)

		// 单次接口结果
		apiResChan := make(chan DecoratorRetryApiResult, 1)

		fmt.Println("[request ] start")
		go func(apiResChan chan DecoratorRetryApiResult) {
			res, err := apiFunc()
			apiResChan <- DecoratorRetryApiResult{res, err}

			// 错误处理
			if err == nil {
				done <- true
				return
			} else {
				retry <- true
			}
		}(apiResChan)
		fmt.Println("[request ] end")

		select {
		// 正常结束
		case <-done:
			decoratorResultChan <- <-apiResChan
			fmt.Println("[response] done!")
			return
		// 错误重试
		case <-retry:
			fmt.Println("[response] error")

			// 比较已经重试次数 是否超过 重试最大次数
			if i >= countRetry {
				fmt.Println("[request ] max retries...")
				decoratorResultChan <- DecoratorRetryApiResult{"", errors.New("response max retries")}
				return
			}
		// 超时重试
		case <-timeAfter: // 单次超时重试
			fmt.Println("[response] timeout")

			// 比较已经重试次数 是否超过 重试最大次数
			if i >= countRetry {
				fmt.Println("[request ] max retries...")
				decoratorResultChan <- DecoratorRetryApiResult{"", errors.New("response max retries")}
				return
			}
		}
	}
}

func main() {
	//_, _ = doErrorApiRequest()

	countRetry := 3      // 重试次数（总请求数 = 重试次数 + 1）
	timeoutResponse := 5 // 响应超时时间（单次请求）
	resultChan := make(chan DecoratorRetryApiResult, 3)

	go retryApiDecorator(doErrorApiRequest, resultChan, countRetry, timeoutResponse)

	select {
	case decoratorResult := <-resultChan:
		fmt.Println("[response]", "data:", decoratorResult.ApiData, "; error:", decoratorResult.ApiError)
		return
	}
}