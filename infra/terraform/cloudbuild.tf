# 接続周りの設定が面倒なので一旦手動で作成した
# TODO: terraform化する
#
# resource "google_cloudbuild_trigger" "learn_golang" {
#   name = "learn-golang-trigger"
#
#   github {
#     owner = "taka011002"
#     name  = "learn-golang"
#     push {
#       branch = "main"
#     }
#   }
#
#   filename = "cloudbuild.yaml"
# }