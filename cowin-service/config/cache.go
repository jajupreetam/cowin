package config

import (
	"bytes"
	"context"
	slog "github.com/go-eden/slf4go"
	"github.com/go-redis/redis/v8"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

const (
	HybrisProduct                 = "hybris-product"
	HybrisBundleTemplate          = "hybris-bundle-template"
	HybrisProductIds              = "hybris-product-ids"
	NuteLineCheckForNewActivation = "nute-linecheck-new-activation"
	NuteLineCheckForModification  = "nute-linecheck-modification"
	OAuthPublicKey                = "oauth-publickey"
	Account                       = "account"
	AccountService                = "account-service"
	DevicePlanEligibility         = "deviceplan-eligibility"
	DevicePlans                   = "deviceplans"
	Migrations                    = "migrations"
	PackMigrations                = "packMigrations"
	PriceOffers                   = "priceOffers"
	Fairpay                       = "fairpay"
	DefaultCacheExpiryInSecond    = 600
)

var Cache cacheManager

type cacheManager struct {
	RedisClient *redis.Client
}

func (cm *cacheManager) Get(key string, obj interface{}) error {
	val, err := Cache.RedisClient.Get(context.Background(), key).Result()
	if err == nil {
		ffjson.Unmarshal([]byte(val), obj)
	}
	return err
}

func (cm *cacheManager) GenerateCacheKey(fun string, params ...string) string {
	var b bytes.Buffer
	b.WriteString(fun)
	b.WriteString("-")
	for i, p := range params {
		str := strings.Replace(strings.TrimSpace(p), " ", "-", -1)
		b.WriteString(str)
		if i < len(params)-1 {
			b.WriteString("_")
		}
	}
	return b.String()
}

func (cm *cacheManager) Set(key string, obj interface{}) (string, error) {
	timeout := viper.GetDuration("cache.timeout")
	return cm.SetWithExpiry(key, obj, time.Second*timeout)
}

func (cm *cacheManager) SetWithExpiry(key string, obj interface{}, expirationInSecond time.Duration) (string, error) {
	marshaled, _ := ffjson.Marshal(obj)
	return Cache.RedisClient.Set(context.Background(), key, marshaled, time.Second*expirationInSecond).Result()
}

func (cm *cacheManager) InitRedisClient() {
	redisUrl := viper.GetString("cache.redis.url")
	redisPassword := os.Getenv("redis_database-password")
	redisDb := viper.GetInt("cache.redis.db")
	slog.Info("Redis Url ", redisUrl)
	slog.Info("Redis Password ", redisPassword)
	slog.Info("Redis Database ", redisDb)
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword,
		DB:       redisDb,
	})
	pong, err := RedisClient.Ping(context.Background()).Result()
	if pong != "PONG" || err != nil {
		slog.Warn("Could not initialize redis ", err)
	} else {
		slog.Info("Successfully initialized redis cache")
	}
	Cache.RedisClient = RedisClient
}
