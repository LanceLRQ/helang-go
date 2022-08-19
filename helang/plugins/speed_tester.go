package plugins

import (
	"fmt"
	"github.com/redmask-hb/GoSimplePrint/goPrint"
	"math"
	"math/rand"
	"time"
)

var MusicList = []string{
	"Kill You", "Lighters", "ZOOD", "Love the Way You Lie",
	"The Monster", "Numb Encore", "Kinds Never Die", "I Need a Doctor",
	"Lose Yourself", "Mockingbird", "Beautiful", "Not Afraid",
	"Rap God", "Phenomenal", "Stan", "Space Bound", "Stan",
	"Guts Over Fear", "Spade",
}

var FileSuffix = []string {
	".mp3", ".flac", ".ogg",
}

var SummaryText = "=== HeLang protects your every pure download ===\nData used | %d MB\nDownloaded files | %d\nLocation | BUPT Xitucheng Campus\n\n=== HeLang protects your pure disk memory ===\nAll test files were deleted\nData freed | %dMB\nDeleted files | %d\nLocation | BUPT Xitucheng Campus\n\n5G speed test finished.\n"

func RunSpeedTest() {
	fmt.Println("Cyber DJ is downloading musics via 5G...")
	rand.Shuffle(len(MusicList), func(i, j int) {
		MusicList[i], MusicList[j] = MusicList[j], MusicList[i]
	})

	totalSize := 0
	for _, music := range MusicList {
		fileSize := 114 + rand.Intn(514)        // ♂, unit MB
		file := music + FileSuffix[rand.Intn(3)]
		vipSuffix := "[VIP]"
		if music == "ZOOD" {
			vipSuffix = ""
		}
		desc := fmt.Sprintf("Downloading %s(%d MB)...%s", file, fileSize, vipSuffix)
		downloadFile(desc, fileSize)
		totalSize += fileSize
	}
	fmt.Println()
	fmt.Printf(SummaryText, totalSize, len(MusicList), totalSize, len(MusicList))
}

func downloadFile(desc string, fileSize int) {
	startTime := time.Now()
	bar := goPrint.NewBar(100)
	bar.SetNotice(desc)
	bar.SetGraph(">")
	c := 0
	avgSpeed := 0.0
	for i := 0; i <= fileSize; i += rand.Intn(50)  {
		bar.PrintBar(int(math.Floor(float64(i) / float64(fileSize) * 100)))
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		nowTime := time.Now()
		t := nowTime.Sub(startTime)
		useTime := float64(t.Milliseconds())
		if useTime > 0 {
			transferSize := float64(i - c)
			speed := transferSize / (useTime / 100)
			if avgSpeed == 0 {
				avgSpeed = speed
			}
			avgSpeed = (speed + avgSpeed) / 2
			c = i
			startTime = time.Now()
		}
	}
	bar.PrintBar(100)
	bar.PrintEnd(fmt.Sprintf("Finish! (%.2f MB/s)", avgSpeed))
}