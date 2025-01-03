# Shopping Notifier

日用品の買い忘れを防ぐためのLIFFアプリです。LINE通知機能を活用し、必要な日用品の在庫状況や購入予定を効率的に管理します。

---

## 特徴

- **日用品リストの管理**: 家庭ごとにリストを作成し、在庫や優先度を共有
- **買い忘れ防止通知**: LINEを通じて必要なタイミングで通知を送信
- **使いやすいUI**: LIFF (LINE Front-end Framework) を活用したシンプルで直感的な操作
- **拡張性**: MySQLとGo言語を基盤に設計されたスケーラブルなアーキテクチャ

---

## 主な機能

1. **ユーザー認証**:
   - LINEログインを使用してユーザーを識別
2. **日用品リストの作成と編集**:
   - 在庫状況、優先度を設定可能
3. **通知機能**:
   - 在庫切れや期限に応じてリマインダーを送信
4. **グループ管理**:
   - 家庭ごとのグループでリストを共有

---

## 技術スタック

- **バックエンド**: Go
- **フロントエンド**: LIFF (LINE Front-end Framework)
- **データベース**: MySQL
- **通知システム**: LINE Messaging API
- **インフラ**: AWS (EC2, RDS)

---

## セットアップ手順

### 1. 必要な環境
- Go >= 1.20
- MySQL >= 8.0
- LINE Developer Account
- AWS Account

### 2. クローンリポジトリ
```bash
git clone https://github.com/your-username/shopping-notifier.git
cd shopping-notifier
