// @Author: wangdehong
// @Description:
// @File: conf
// @Date: 2022/4/23 7:16 下午

package conf

// Config is comet config.
type Config struct {
	Debug     bool
	Env       *Env
	Bucket    *Bucket
}

// Env is env config.
type Env struct {
	Region    string
	Zone      string
	DeployEnv string
	Host      string
	Weight    int64
	Offline   bool
	Addrs     []string
}

// Bucket is bucket config.
type Bucket struct {
	Size          int
	Channel       int
	Room          int
	RoutineAmount uint64
	RoutineSize   int
}
