package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/winterssy/instabot"
	"madara.io/test/insta/inst"
)

var (
	user       string
	password   string
	target     string
	like       int
	skipped    int
	days       int
	errorCount int
	bot        *instabot.Bot
	ctx        context.Context
	cancel     context.CancelFunc
	followers  []inst.User
)

func init() {
	rand.Seed(time.Now().Unix())

	flag.StringVar(&user, "u", "", "Your Instagram username")
	flag.StringVar(&password, "p", "", "Your Instagram password")
	flag.StringVar(&target, "t", "", "Target to fetch this followers")
	flag.IntVar(&like, "l", 5, "Number of items to like on followers target")
	flag.IntVar(&days, "d", 90, "Number of days to skip posts older than this number")
	flag.Parse()

	if user == "" || password == "" || target == "" {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	bot = instabot.New(user, password)
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	_, err := bot.Login(ctx, true)
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}

	var user inst.GetUserByName
	resp, err := bot.GetUserByName(ctx, target)
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
	resp.Decode(&user)

	getUserFollowers(user.User, "")
	rand.Shuffle(len(followers), func(i, j int) { followers[i], followers[j] = followers[j], followers[i] })
	for k, f := range followers {
		color.Cyan("🌅Fetch feed from %s %v/%v\n", f.Username, k+1, len(followers))
		var feed inst.GetUserFeed
		resp, err = bot.GetUserFeed(ctx, fmt.Sprintf("%v", f.Pk), "")
		if err != nil {
			color.Red(err.Error())
			os.Exit(0)
		}
		resp.Decode(&feed)
		likeItems(like, feed.Items)
	}
	color.Yellow("❌Skipped %v of %v\n", skipped, len(followers))
	color.Cyan("✅Like done !\n")

}

func getUserFollowers(u inst.User, maxID string) {
	if len(followers) == 0 {
		color.Magenta("👥Fetch %d followers from %s\n", u.FollowerCount, u.Username)
	}
	var uFollowers inst.GetUserFollowers
	resp, err := bot.GetUserFollowers(ctx, fmt.Sprintf("%v", u.Pk), fmt.Sprintf("%v", maxID))
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
	resp.Decode(&uFollowers)

	for _, uf := range uFollowers.Users {
		if !uf.IsPrivate {
			followers = append(followers, uf)
		}
	}
	color.Green("👥Append %v publics followers of %v from %s\n", len(followers), u.FollowerCount, u.Username)

	if len(uFollowers.NextMaxID) > 1 {
		color.Cyan("⏩Paginate next followers from %s\n", u.Username)
		sleepTime(5)
		getUserFollowers(u, uFollowers.NextMaxID)
	}
}

func likeItems(maxLike int, items []inst.Item) {
	l := 1

	if len(items) == 0 {
		skipped++
		color.Yellow("🤷‍♀️Nothing to like skip %v of %v...\n", skipped, len(followers))
	}

	for _, i := range items {
		if !i.HasLiked {
			if !isToOld(i) {
				if l <= maxLike {
					color.Green("💚Like %v/%v\n", l, maxLike)
					_, err := bot.LikeMedia(ctx, i.ID, false)
					if err != nil {
						color.Red(err.Error())
						errorCount++
						if errorCount >= 20 {
							os.Exit(0)
						}
					}
					l++
					sleepTime(15)
				} else {
					break
				}
			} else {
				skipped++
				color.Yellow("🧓Item are too old skip %v of %v...\n", skipped, len(followers))
				break
			}
		} else {
			skipped++
			color.Yellow("😅Already liked skip %v of %v...\n", skipped, len(followers))
			break
		}
	}
}

func sleepTime(max int) {
	min := 1
	ran := rand.Intn(max-min) + min
	color.Magenta("😴Sleep %vsec...\n", ran)
	time.Sleep(time.Duration(int64(ran)) * time.Second)
}

func isToOld(i inst.Item) bool {
	tm := time.Unix(int64(i.Caption.CreatedAt), 0)
	n := time.Now()
	d := n.Sub(tm).Hours() / 24
	if int(d) > days {
		return true
	}
	return false
}
