# go_blockingqueue

sync.mutex를 이용하여 LinkedBlockingQueue를 구현해 보았다.

## 큐생성

```go
var capacity uint64
capacity = 100
queue, err := blockingQueue.NewBlockingQueue(capacity)
if err != nil {
  t.Fatal(err.Error())
}
```
capacity만큼 list 생성

### Push

```go
for i := uint64(0); i < capacity+1; i++ {
  _, err := queue.Push(i)
  if err != nil {
    t.Log(err.Error())
  } else {

  }
}
```
capacity 초과할시 err가 반환된다.

## Pop

```go
for i := uint64(0); i < capacity; i++ {
  _, err := queue.Pop()
  if err != nil {
    t.Log(err.Error())
  } else {

  }
}
```
큐가 비었을 경우 Pop()에서 에러가 발생한다.

> [!note]
> 
> 에러가 나는게 정상인지 의심스럽다만 일단은 에러로 처리하고 추후 polling이 필요할 듯하다.

### Go 루틴
```go
var groupCount int
groupCount = 5
wg := new(sync.WaitGroup)
wg.Add(groupCount)

for i := 0; i < groupCount; i++ {
  go func(n int) {
    for queue.IsEmpty() != true {
      res, err := queue.Pop()
      if err == nil {
        t.Logf("pop[%d] : %d\n", n+1, res)
        time.Sleep(1 * time.Second)
      } else {

      }
    }
    wg.Done()
  }(i)
}

wg.Wait()
fmt.Println("done...")
```
mutex가 잘 적용 된 듯하다.
