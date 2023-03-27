/*
data "ochk_backup_plan" "backup_plan" {
  display_name = var.backup_plan
}
*/

/*
data "ochk_backup_list" "backup_list" {
  display_name = var.backup_list
  backup_plan_id = data.ochk_backup_plan.backup_plan.id
}
*/

data "ochk_backup_plans" "backup_plans" {
}

/*
data "ochk_backup_lists" "backup_lists" {
  //backup_plan_id = data.ochk_backup_plans.backup_plans.backup_plans[0].backup_plan_id
  backup_plan_id = data.ochk_backup_plan.backup_plan.id
}
*/