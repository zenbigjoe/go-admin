// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package language

import "strings"

var jp = LangSet{
	"managers":  "管理者",
	"name":      "ユーザー名",
	"nickname":  "ニックネーム",
	"role":      "ロール",
	"createdat": "作成日時",
	"updatedat": "更新日時",
	"path":      "パス",
	"submit":    "提出",
	"filter":    "フィルター",

	"new":             "新規",
	"export":          "出力",
	"action":          "操作",
	"toggle dropdown": "プルダウン",
	"delete":          "削除",
	"refresh":         "更新",
	"back":            "戻る",
	"reset":           "リセット",
	"save":            "保存",
	"edit":            "編集",
	"expand":          "展開",
	"collapse":        "折り畳み",
	"online":          "オンライン",
	"setting":         "設定",
	"sign out":        "ログアウト",
	"home":            "ホーム",
	"all":             "全て",
	"more":            "更に",
	"remove":          "削除する",

	"permission manage": "権限管理",
	"menus manage":      "メニュー管理",
	"roles manage":      "ロール管理",
	"operation log":     "操作履歴",
	"method":            "メソッド",
	"input":             "入力",
	"operation":         "操作",
	"menu name":         "表示名",
	"reload succeeded":  "再読み込み完了",
	"search":            "検索",

	"permission denied": "権限がありません",
	"error":             "エラー",
	"current page":      "現在のページ",
	"query time":        "クエリ時間",

	"are you sure to delete": "本当に削除しますか",
	"yes":                    "はい",
	"got it":                 "はい",
	"cancel":                 "キャンセル",
	"refresh succeeded":      "更新に成功しました",
	"edit fail":              "編集に失敗しました",
	"create fail":            "作成に失敗しました",
	"confirm password":       "パスワード(確認)",
	"all method if empty":    "空の場合は全メソッド",

	"detail": "詳細",

	"avatar":     "アバター",
	"password":   "パスワード",
	"username":   "ユーザー名",
	"slug":       "スラッグ",
	"permission": "権限",
	"userid":     "ユーザーID",
	"content":    "内容",
	"parent":     "親メニュー",
	"icon":       "アイコン",
	"uri":        "URL",

	"login":      "ログイン",
	"login fail": "ログイン失敗",

	"admin":     "運営管理",
	"users":     "ユーザー",
	"roles":     "ロール",
	"menu":      "メニュー",
	"dashboard": "ダッシュボード",

	"goadmin is now running. \nrunning in \"debug\" mode. switch to \"release\" mode in production.\n\n": "GoAdminが現在稼働中です。\n\"デバッグ\"モードで実行中です。本番環境では\"リリース\"モードに切り替えてください。\n\n",

	"wrong goadmin version, theme %s required goadmin version are %s":    "GoAdminバージョンが間違っています、テーマ%sの必要なGoAdminバージョンは%sです。",
	"wrong theme version, goadmin %s required version of theme %s is %s": "テーマのバージョンが間違っています、GoAdmin%sの必要なテーマ%sのバージョンは%sです。",

	"adapter is nil, import the default adapter or use addadapter method add the adapter": "アダプターがnilです。デフォルトのアダプターをインポートするか、AddAdapterメソッドを使用してアダプターを追加してください。",

	"username and password can not be empty":        "アカウントまたパスワードが正しく入力されていることを確認してください",
	"operation not allow":                           "この操作を実行するアクセス許可が必要です",
	"password does not match":                       "パスワードが正しくありません",
	"should be unique":                              "重複しないことを確認してください",
	"slug exists":                                   "スラッグが既に存在していないことを確認してください",
	"no corresponding options?":                     "対応できるオプションがありません",
	"create here.":                                  "ここに新規作成",
	"use for login":                                 "ログインに使う",
	"use to display":                                "表示に使う",
	"a path a line, without global prefix":          "パスを１行ずつ入力してください",
	"slug or http_path or name should not be empty": "スラッグ、http_pathまたユーザー名が正しく入力されていることを確認してください",
	"no roles":                                      "ロールなし",

	"initialize configuration":        "設定の初期化",
	"initialize navigation buttons":   "ナビゲーションボタンの初期化",
	"initialize plugins":              "プラグインの初期化",
	"initialize database connections": "データベース接続の初期化",
	"initialize success":              "初期化成功🍺🍺 ",

	"not found":      "見つかりません",
	"internal error": "内部エラー",
	"unauthorized":   "権限がありません",

	"plugins":          "プラグイン",
	"plugin store":     "プラグインストア",
	"get more plugins": "もっとプラグインを取得する",
	"uninstalled":      "アンインストールされました",
	"plugin setting":   "プラグイン設定",

	"second":  "秒",
	"seconds": "秒",
	"minute":  "分",
	"minutes": "分",
	"hour":    "時間",
	"hours":   "時間",
	"day":     "日",
	"days":    "日",
	"week":    "週間",
	"weeks":   "週間",
	"month":   "月",
	"months":  "月",
	"year":    "年",
	"years":   "年",

	"config.domain":          "ウェブサイトのドメイン",
	"config.language":        "ウェブサイトの言語",
	"config.url prefix":      "URLプレフィックス",
	"config.theme":           "テーマ",
	"config.title":           "タイトル",
	"config.index url":       "ホームURL",
	"config.login url":       "ログインURL",
	"config.env":             "環境",
	"config.color scheme":    "カラースキーム",
	"config.cdn url":         "CDNアセットURL",
	"config.login title":     "ログインタイトル",
	"config.auth user table": "認証ユーザーテーブル",
	"config.extra":           "追加の設定",
	"config.store":           "ファイルストアの設定",
	"config.databases":       "データベースの設定",

	"config.general":      "一般",
	"config.log":          "ログ",
	"config.site setting": "サイト設定",
	"config.custom":       "カスタマイズ",
	"config.debug":        "デバッグモード",
	"config.site off":     "サイトオフライン",
	"config.true":         "オン",
	"config.false":        "オフ",

	"config.logo":                        "ロゴ",
	"config.mini logo":                   "ミニロゴ",
	"config.bootstrap file path":         "Bootstrapファイルのパス",
	"config.go mod file path":            "go.modファイルのパス",
	"config.session life time":           "セッションの有効期間",
	"config.custom head html":            "ヘッドHTML",
	"config.custom foot html":            "フッターHTML",
	"config.custom 404 html":             "404ページ",
	"config.custom 403 html":             "403ページ",
	"config.custom 500 html":             "500ページ",
	"config.hide config center entrance": "設定ボタンを非表示にする",
	"config.hide app info entrance":      "アプリ情報ボタンを非表示にする",
	"config.hide tool entrance":          "ツールボタンを非表示にする",
	"config.hide plugin entrance":        "プラグインリストボタンを非表示にする",
	"config.footer info":                 "フッター情報",
	"config.login logo":                  "ログインロゴ",
	"config.no limit login ip":           "制限なしの複数IPでのログイン",
	"config.operation log off":           "操作ログオフ",
	"config.allow delete operation log":  "操作ログの削除を許可する",
	"config.animation type":              "アニメーションタイプ",
	"config.animation duration":          "アニメーションの期間（秒）",
	"config.animation delay":             "アニメーションの遅延（秒）",
	"config.file upload engine":          "ファイルアップロードエンジン",

	"config.logger rotate":             "ログローテーション設定",
	"config.logger rotate max size":    "最大サイズ（MB）",
	"config.logger rotate max backups": "最大バックアップ数",
	"config.logger rotate max age":     "最大保存期間（日）",
	"config.logger rotate compress":    "圧縮",

	"config.info log path":         "情報ログファイルのパス",
	"config.error log path":        "エラーログファイルのパス",
	"config.access log path":       "アクセスログファイルのパス",
	"config.info log off":          "情報ログオフ",
	"config.error log off":         "エラーログオフ",
	"config.access log off":        "アクセスログオフ",
	"config.access assets log off": "アクセスアセットログオフ",
	"config.sql log on":            "SQLログを開く",
	"config.log level":             "レベル",

	"config.logger rotate encoder":                "ログエンコーダー設定",
	"config.logger rotate encoder time key":       "時間キー",
	"config.logger rotate encoder level key":      "レベルキー",
	"config.logger rotate encoder name key":       "名前キー",
	"config.logger rotate encoder caller key":     "呼び出し元キー",
	"config.logger rotate encoder message key":    "メッセージキー",
	"config.logger rotate encoder stacktrace key": "スタックトレースキー",
	"config.logger rotate encoder level":          "レベルエンコーダー",
	"config.logger rotate encoder time":           "時間エンコーダー",
	"config.logger rotate encoder duration":       "持続時間エンコーダー",
	"config.logger rotate encoder caller":         "呼び出し元エンコーダー",
	"config.logger rotate encoder encoding":       "出力形式",

	"config.capital":        "大文字",
	"config.capitalcolor":   "色付き大文字",
	"config.lowercase":      "小文字",
	"config.lowercasecolor": "色付き小文字",

	"config.seconds":     "秒",
	"config.nanosecond":  "ナノ秒",
	"config.microsecond": "マイクロ秒",
	"config.millisecond": "ミリ秒",

	"config.full path":  "完全なパス",
	"config.short path": "短縮パス",

	"config.do not modify when you have not set up all assets": "すべてのアセットが設定されていない場合は変更しないでください",
	"config.it will work when theme is adminlte":               "テーマがadminlteの場合に機能します",

	"config.language." + CN:                  "中国語",
	"config.language." + EN:                  "英語",
	"config.language." + JP:                  "日本語",
	"config.language." + PTBR:                "ブラジルポルトガル語",
	"config.language." + strings.ToLower(TC): "繁体字中国語",
	"config.language." + RU:                  "ロシア語",

	"config.modify site config":         "サイト設定の変更",
	"config.modify site config success": "変更成功",
	"config.modify site config fail":    "変更失敗",

	"system.system info":     "システムおよびアプリケーション情報",
	"system.application":     "アプリケーション情報",
	"system.application run": "実行中アプリケーション情報",
	"system.system":          "システム情報",

	"system.process_id":                           "プロセスID",
	"system.golang_version":                       "Golangバージョン",
	"system.server_uptime":                        "サーバーの稼働時間",
	"system.current_goroutine":                    "現在のゴールーチン数",
	"system.current_memory_usage":                 "現在のメモリ使用量",
	"system.total_memory_allocated":               "割り当てられた総メモリ量",
	"system.memory_obtained":                      "取得されたメモリ",
	"system.pointer_lookup_times":                 "ポインター検索回数",
	"system.memory_allocate_times":                "メモリ割り当て回数",
	"system.memory_free_times":                    "メモリ解放回数",
	"system.current_heap_usage":                   "現在のヒープ使用量",
	"system.heap_memory_obtained":                 "取得されたヒープメモリ",
	"system.heap_memory_idle":                     "アイドル状態のヒープメモリ",
	"system.heap_memory_in_use":                   "使用中のヒープメモリ",
	"system.heap_memory_released":                 "解放されたヒープメモリ",
	"system.heap_objects":                         "ヒープオブジェクト数",
	"system.bootstrap_stack_usage":                "ブートストラップスタック使用量",
	"system.stack_memory_obtained":                "取得されたスタックメモリ",
	"system.mspan_structures_usage":               "MSpan構造の使用量",
	"system.mspan_structures_obtained":            "取得されたMSpan構造",
	"system.mcache_structures_usage":              "MCache構造の使用量",
	"system.mcache_structures_obtained":           "取得されたMCache構造",
	"system.profiling_bucket_hash_table_obtained": "プロファイリングバケットハッシュテーブルの取得",
	"system.gc_metadata_obtained":                 "GCメタデータの取得",
	"system.other_system_allocation_obtained":     "その他のシステム割り当ての取得",
	"system.next_gc_recycle":                      "次のGC再利用",
	"system.last_gc_time":                         "最後のGCからの時間経過",
	"system.total_gc_time":                        "総GC時間",
	"system.total_gc_pause":                       "総GC停止時間",
	"system.last_gc_pause":                        "最後のGC停止時間",
	"system.gc_times":                             "GC回数",

	"system.cpu_logical_core": "CPU論理コア",
	"system.cpu_core":         "CPU物理コア",
	"system.os_platform":      "OSプラットフォーム",
	"system.os_family":        "OSファミリー",
	"system.os_version":       "OSバージョン",
	"system.load1":            "Load1",
	"system.load5":            "Load5",
	"system.load15":           "Load15",
	"system.mem_total":        "総メモリ量",
	"system.mem_available":    "利用可能なメモリ",
	"system.mem_used":         "使用中のメモリ",

	"system.app_name":         "アプリ名",
	"system.go_admin_version": "アプリバージョン",
	"system.theme_name":       "テーマ",
	"system.theme_version":    "テーマバージョン",

	"tool.tool":                   "ツール",
	"tool.table":                  "テーブル",
	"tool.connection":             "接続",
	"tool.package":                "パッケージ",
	"tool.output":                 "出力パス",
	"tool.output path is empty":   "出力パスが空です",
	"tool.field":                  "データベースフィールド",
	"tool.title":                  "タイトル",
	"tool.field name":             "名前",
	"tool.db type":                "データベースタイプ",
	"tool.form type":              "フォームタイプ",
	"tool.generate table model":   "テーブルモデルを生成する",
	"tool.primarykey":             "主キー",
	"tool.field filterable":       "フィルター可能",
	"tool.field sortable":         "ソート可能",
	"tool.yes":                    "はい",
	"tool.no":                     "いいえ",
	"tool.hide":                   "非表示",
	"tool.show":                   "表示",
	"tool.generate success":       "生成成功",
	"tool.hide filter area":       "フィルターエリアを非表示にする",
	"tool.use absolute path":      "絶対パスを使用する",
	"tool.display":                "表示",
	"tool.basic info":             "基本",
	"tool.table info":             "テーブル",
	"tool.form info":              "フォーム",
	"tool.field editable":         "編集可能",
	"tool.info field editable":    "編集非表示",
	"tool.field can add":          "追加可能",
	"tool.field default":          "デフォルト",
	"tool.filter area":            "フィルターエリア",
	"tool.new button":             "新規ボタン",
	"tool.extra import package":   "パッケージのインポート",
	"tool.export button":          "エクスポートボタン",
	"tool.edit button":            "編集ボタン",
	"tool.delete button":          "削除ボタン",
	"tool.detail button":          "詳細ボタン",
	"tool.filter button":          "フィルターボタン",
	"tool.row selector":           "行選択器",
	"tool.pagination":             "ページネーション",
	"tool.query info":             "クエリ情報",
	"tool.filter form layout":     "フィルターフォームレイアウト",
	"tool.continue edit checkbox": "編集の続行チェックボックス",
	"tool.continue new checkbox":  "新規作成の続行チェックボックス",
	"tool.reset button":           "リセットボタン",
	"tool.back button":            "戻るボタン",
	"tool.generate":               "生成する",
	"tool.generated tables":       "生成されたテーブル",
	"tool.description":            "説明",
	"tool.label":                  "ラベル",
	"tool.image":                  "画像",
	"tool.bool":                   "ブール",
	"tool.link":                   "リンク",
	"tool.fileSize":               "ファイルサイズ",
	"tool.date":                   "日付",
	"tool.icon":                   "アイコン",
	"tool.dot":                    "ドット",
	"tool.progressBar":            "プログレスバー",
	"tool.loading":                "ロード中",
	"tool.downLoadable":           "ダウンロード可能",
	"tool.copyable":               "コピー可能",
	"tool.carousel":               "カルーセル",
	"tool.qrcode":                 "QRコード",
	"tool.field hide":             "非表示",
	"tool.field display":          "表示",
	"tool.table permission":       "権限を生成する",
	"tool.extra code":             "追加コード",

	"tool.field display normal":     "通常",
	"tool.field diplay hide":        "非表示",
	"tool.field diplay edit hide":   "編集非表示",
	"tool.field diplay create hide": "作成非表示",

	"tool.detail display":             "表示",
	"tool.detail info":                "詳細情報",
	"tool.follow list page":           "リストページに従う",
	"tool.inherit from list page":     "リストページから継承",
	"tool.independent from list page": "リストページから独立",

	"tool.generate table model success": "生成成功",
	"tool.generate table model fail":    "生成失敗",

	"generator.query":                 "クエリ",
	"generator.show edit form page":   "編集フォームページを表示",
	"generator.show create form page": "作成フォームページを表示",
	"generator.edit":                  "編集",
	"generator.create":                "作成",
	"generator.delete":                "削除",
	"generator.export":                "エクスポート",

	"plugin.plugin":                         "プラグイン",
	"plugin.plugin detail":                  "プラグイン詳細",
	"plugin.introduction":                   "導入",
	"plugin.website":                        "ウェブサイト",
	"plugin.version":                        "バージョン",
	"plugin.created at":                     "作成日時",
	"plugin.updated at":                     "更新日時",
	"plugin.provided by %s":                 "%sに提供されます",
	"plugin.upgrade":                        "アップグレード",
	"plugin.install":                        "インストール",
	"plugin.info":                           "詳細",
	"plugin.download":                       "ダウンロード",
	"plugin.buy":                            "購入",
	"plugin.downloading":                    "ダウンロード中",
	"plugin.login":                          "ログイン",
	"plugin.login to goadmin member system": "GoAdminメンバーシステムにログインする",
	"plugin.account":                        "アカウント",
	"plugin.password":                       "パスワード",
	"plugin.learn more":                     "もっと学ぶ",

	"plugin.no account? click %s here %s to register.":    "アカウントがありませんか？登録するには%sここ%sをクリックしてください。",
	"plugin.download fail, wrong name":                    "ダウンロードに失敗しました、間違った名前です",
	"plugin.change to debug mode first":                   "まずデバッグモードに変更してください",
	"plugin.download fail, plugin not exist":              "ダウンロードに失敗しました、プラグインが存在しません",
	"plugin.download fail":                                "ダウンロードに失敗しました",
	"plugin.golang develop environment does not exist":    "Golangの開発環境が存在しません",
	"plugin.download success, restart to install":         "ダウンロードが成功しました、インストールのために再起動します",
	"plugin.restart to install":                           "インストールのために再起動します",
	"plugin.can not connect to the goadmin remote server": "GoAdminリモートサーバーに接続できませんでした、ネットワーク接続を確認してください。",

	"admin.basic admin": "基本管理者",
	"admin.a built-in plugins of goadmin which help you to build a crud manager platform quickly.": "GoAdminの組み込みプラグインで、CRUDマネージャープラットフォームを素早く構築できます。",
	"admin.official": "公式",

	"browse":                              "開く",
	"system.site info":                    "実行情報",
	"delete succeed":                      "削除成功",
	"continue creating":                   "追加続行",
	"wrong captcha":                       "エラーの認証コード",
	"code generate tool":                  "コードジェネレーター",
	"fixed the sidebar":                   "固定サイドバー",
	"config.must bigger than 900 seconds": "900秒以上である必要があります",
	"enter fullscreen":                    "フルスクリーンに入る",
	"fail":                                "失敗",
	"modify success":                      "修正成功",
	"menus":                               "メニュー",
	"config.test":                         "テスト環境",
	"del":                                 "削除",
	"close":                               "閉じる",
	"view":                                "表示",
	"tool.code generate tool":             "コードジェネレーター",
	"login info":                          "ログイン情報",
	"login overdue, please login again":   "ログイン情報が期限切れです。再ログインしてください。",
	"site setting":                        "ウェブサイト設定",
	"managers manage":                     "管理者管理",
	"continue editing":                    "編集を続ける",
	"site info":                           "実行情報",
	"success":                             "成功",
	"delete fail":                         "削除失敗",
	"exit fullscreen":                     "フルスクリーンを終了する",
	"confirm":                             "確認",
	"user":                                "ユーザー",
	"showing <b>%s</b> to <b>%s</b> of <b>%s</b> entries": "第<b>%s</b>から第<b>%s</b>のレコードを表示、合計<b>%s</b>レコード",
	"config.local":  "ローカル環境",
	"config.prod":   "プロダクション環境",
	"no permission": "権限がありません",
}
