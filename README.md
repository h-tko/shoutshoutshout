# このアプリについて

このアプリはGoからWatson Conversationを呼び出すサンプルアプリです。

# 利用方法

このアプリはwebpackとnpmを使用していますので、git clone後は以下を実行してください。

```
$ npm install

$ webpack
```

libraries/watson.goの定数を、ご自身のConversationWorkspaceの登録情報に置き換えてください。

```go:libraries/watson.go
const (
    BLUEMIX_CONVERSATION_URL  string = "https://gateway.watsonplatform.net/conversation/api/v1/workspaces/%s/message?version=%s"
    CONVERSATION_WORKSPACE_ID string = "xxxxxxxxxxxxx"  //　この辺
    CONVERSATION_USERNAME     string = "xxxxxxxxxxxxx" //　この辺
    CONVERSATION_PASSWORD     string = "xxxxxxxxxxxx" //　この辺
    VERSION                   string = "yyyy-mm-dd" //　この辺
)
```

また、このサンプルアプリはgo_buildpackを利用してbluemixにデプロイするようにできていますので、デプロイ時は以下コマンドでデプロイしてください。

```
$ cf push <application_name> -b go_buildpack
```
