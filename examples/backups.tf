data "ochk_backup_plan" "backup_plan" {
  display_name = var.backup_plan
}

data "ochk_backup_list" "backup_list" {
  display_name = var.backup_list
  backup_plan_id = data.ochk_backup_plan.backup_plan.id
}