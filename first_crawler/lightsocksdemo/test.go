// const PasswordLength = 256
// type Password [PasswordLength]byte

// func init() {
// 	// 更新随机种子，防止生成一样的随机密码
// 	rand.Seed(time.Now().Unix())
// }

// // 产生 256个byte随机组合的 密码
// func RandPassword() *Password {
// 	// 随机生成一个由  0~255 组成的 byte 数组
// 	intArr := rand.Perm(PasswordLength)
// 	password := &Password{}
// 	for i, v := range intArr {
// 		password[i] = byte(v)
// 		if i == v {
// 			// 确保不会出现如何一个byte位出现重复
// 			return RandPassword()
// 		}
// 	}
// 	return password
// }
package main
import (
	"context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
)

func main(){
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	imageList, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	for x :=range imageList{
		println(x)
	}

}