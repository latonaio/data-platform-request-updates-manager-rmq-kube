# data-platform-request-updates-manager-rmq-kube
data-platform-request-updates-manager-rmq-kube は、フロントエンドUI等 からの DBデータ登録更新要求に基づいて、バックエンドのcreates/updates/requests/generates系等 マイクロサービス郡 ならびに DB更新機能を担う sql-update-kube とのやり取りをコントロールしながら中継仲介する マイクロサービスです。  

## 動作環境
data-platform-request-updates-manager-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  

## api-input-reader
api-input-reader は、フロントエンドUI等 からのデータ登録更新要求条件を定義するタイプフォーマットです。  

## api-module-runtimes-requests
api-module-runtimes-requests は、フロントエンドUIからのデータ登録更新要求条件をもとに、バックエンドに対してリクエストを要求するファンクションプログラムとパターンを記述します。  

## api-module-runtimes-responses
api-module-runtimes-responses は、バックエンドからのデータのレスポンスを定義するタイプフォーマットです。  

## api-output-formatter
api-output-formatter は、レスポンスされたデータをフロントエンドUI等に返す際のフォーマット定義です。

## controllers
controllers は、本マイクロサービスのコアで、UIからの登録更新要求を、UIの要求に応じて記述します。  
バックエンドの、しばしば複数マイクロサービスとの複雑なデータ登録更新要求が記載されます。　　

## services
services 内の機能は、全controllerに共通の共通サービスを定義する場合に使用されます。  

## routers
routersは、controller の各機能をUI のURL制御と関連つけるための定義を記述します。  