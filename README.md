# mackerel-agent installer for QNAP

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/Songmu/ghch/blob/master/LICENSE

## DESCRIPTION

QNAPにmackerel-agentをインストールして自動起動までやってくれる君です。作者の環境にあわせて、 `GOARM=5` に今のところ決め打ちしてます。

## USAGE

QNAP上でこのrepositoryをcloneなり、アーカイブをダウンロードなりして、 `make install` するだけです。

    % wget https://github.com/Songmu/qnap-mackerel-agent/archive/master.tar.gz
    % tar xzvf master.tar.gz
    % cd qnap-mackerel-agent-master
    % make install
    % /share/MD0_DATA/.mackerel-agent/run.sh start

## INSTALLATION

`make install` 時にいくつかの環境変数を渡してオプション設定が可能です。

### `MACKEREL_API_KEY`
Mackerelの API Keyを設定します。設定されていない場合はプロンプトが表示され入力が促されます。

### `MACKERE_INSTALL_PATH`
mackerel-agent関連ファイルの配置ディレクトリです。デフォルトは `/share/MD0_DATA/.mackerel-agent` です。RAID構成じゃないQNAPの場合、MD0_DATAの部分がMDA_DATAになったりするようなので適宜調整してください。

このインストールパス配下に、mackerel-agentの設定やプラグインなども配置していくことになるため、このパスは覚えておいてください。

## LAUNCH

一般的な起動スクリプト同様に以下でmackerel-agentが起動します。OS起動時には自動起動されるため、このコマンドを手動で実行する必要はありません。

    % /share/MD0_DATA/.mackerel-agent/run.sh start

`start`の他に`stop`, `restart` を渡すこともできます。設定を変更した場合には `restart` するようにしてください。

## UPDATE

新しいバージョンを落として、 `make install` しなおすだけです。

ただし、MackerelのAPIキーについては、一度設定してしまうと `make install` では更新ができません。変更したい場合は、インストールパス直下の `mackerel-agent.conf` を直接編集してください。

## ADVANCED

プラグインを利用するなど、追加の設定を書きたい場合にはインストールパス以下の `conf.d/*.conf` に設定を配置してください。

これは、インストーラーが自動生成する `mackerel-agent.conf` のテンプレートに以下のようなinclude記述があり、そこで指定されています。

```
include = "{{.InstallPath}}/conf.d/*.conf"
```

`mackerel-agent.conf` を直接編集するのではなく、 `conf.d` 配下に、設定ファイルやプラグイン類を配置し、構成管理をすることをオススメします。

## AUTHOR

[Songmu](https://github.com/Songmu)
