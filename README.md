Podcast for onsen
-----------------

[インターネットラジオステーション＜音泉＞](http://www.onsen.ag/)のポッドキャストを生成するスクリプト

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/kosh04/onsenradio-podcast)

## Usage

    $ ./onsenradio-podcast -help
    Usage of ./onsenradio-podcast:
      -addr address
        	HTTP serve address (default ":8000")
      -http
        	Start HTTP serve mode

## ポッドキャストを登録する

    $ ./onsenradio-podcast -http

このプログラムをサーバモードで起動した後
http://localhost:8000/podcast をポッドキャストURLとして登録してください。

また、このプログラムはHeroku上でも稼働しています。面倒な人は次のURLを登録してください。

https://onsenradio.herokuapp.com/podcast

## TODO

- [ ] 更新日時順ソート
- [ ] podcast.Item.Description の定義
- [ ] バナー画像は表示できる？
- [ ] onsen.ag への負荷対策（キャッシュ等）

## WARNING

> ※＜音泉＞アプリリニューアル後は、スマホからの視聴はアプリでのみとなります。

上記の注意書きがあるように、生成元サイトの仕様変更によってコードが
利用できなくなる可能性があります。ご利用は自己責任でお願いします。
