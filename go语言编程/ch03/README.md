# 项目介绍

全部源码在 src 文件夹下，mplayer 里是主程序，mlib 中是对播放器的控制逻辑，mp 中是对播放操作的逻辑。


## 项目初始化
```
go init music 

```

## 执行文件

```
$  go run mplayer.go

                        Enter folloing commands to control the player:
                        lib list --view the existing music lib
                        lib add <name><artist><source><type> --Add a music to the music lib
                        lib remove <name> -- Remove the specified music from the lib
                        play <name> -- Play the specified music

Enter command-> lib add HugeStone MJ pop  ~/MusicLib/hs.mp3 MP3 
USAGE: lib add <name><artist><source><type> 8
Enter command-> lib add HugeStone MJ pop ~/MusicLib/hs.mp3 MP3  
Enter command-> play HugeStone
Playing ~/MusicLib/hs.mp3
Playing ~/MusicLib/hs.mp3 10%
Playing ~/MusicLib/hs.mp3 20%
Playing ~/MusicLib/hs.mp3 30%
Playing ~/MusicLib/hs.mp3 40%
Playing ~/MusicLib/hs.mp3 50%
Playing ~/MusicLib/hs.mp3 60%
Playing ~/MusicLib/hs.mp3 70%
Playing ~/MusicLib/hs.mp3 80%
Playing ~/MusicLib/hs.mp3 90%
Playing ~/MusicLib/hs.mp3 100%

finished Playing ~/MusicLib/hs.mp3 done
Enter command-> lib list
1 : HugeStone pop ~/MusicLib/hs.mp3 MP3
Enter command-> lib add HugeStone2 MJ2 pop ~/MusicLib/hs.mp3 MP3
Enter command-> lib list
1 : HugeStone pop ~/MusicLib/hs.mp3 MP3
2 : HugeStone2 pop ~/MusicLib/hs.mp3 MP3
Enter command-> list view
Unrecognized command: list
Enter command-> q
```