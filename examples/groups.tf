/*
data "ochk_group" "admin-InfraAdm" {
    name="${var.test-data-prefix}-bg-01-InfraAdm"
    depends_on = [
        ochk_subtenant.sub-1
    ]
}
*/

/*
resource "ochk_local_group" "infra-admins" {
    name = "${var.test-data-prefix}-my_infa-admins9"
    users = [
        data.ochk_user.testy-radek.id,
        data.ochk_user.testy-userpb.id
    ]
    parent_groups = [
        data.ochk_group.admin-InfraAdm.id
    ]
    child_groups = [

    ]
}
*/
