<?php
/**
 * Created by PhpStorm.
 * Date: 16-7-13
 * Time: 下午2:13
 */
$oaService = new \app\models\lib\LibOaService();
$this->title = '用户添加';
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 权限管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['user/index']);?>">用户列表</a></li>
            <li class="active"><a href="#">用户添加</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">添加</h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 用户添加</header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo \Yii::$app->urlManager->createUrl(['user/add']);?>" method="post" data-validate="parsley">
                <div class="form-group"><label class="col-lg-2 control-label">用户名：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="wangmoumou" data-required="true" class="form-control parsley-validated"  name="username" data-minlength="3" id="username">
                    <a href="javascript:void(0)" onclick="show_user()" class="btn btn-danger" style="margin-top: 10px">择OA用户</a>
                    </div>
                </div>
                <div class="form-group"><label class="col-lg-2 control-label">姓名：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="某某" data-required="true" class="form-control parsley-validated"  name="nickname" data-minlength="2" id="nickname">
                    </div>
                </div>

                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">密码：</label>
                    <div class="col-sm-10">
                        <input id="pwd" class="form-control parsley-validated" type="password" data-required="true" name="password" placeholder="minlength = 6" data-minlength="6">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">再次输入密码：</label>
                    <div class="col-sm-10">
                        <input class="form-control parsley-validated" type="password" data-required="true" name="confirm_password" data-equalto="#pwd" placeholder="minlength = 6" data-minlength="6">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">角色：</label>

                    <div class="col-sm-10">
                        <select class="form-control parsley-validated" data-required="true" name="role_id">
                            <option value="">请选择</option>
                            <?php foreach ($role as $val): ?>
                                <option value="<?php echo $val['role_id'];?>"><?php echo $val['role_name'];?> </option>
                            <?php endforeach;?>
                        </select>
                    </div>
                </div>

                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">状态：</label>
                    <div class="col-sm-10">
                        <div class="col-sm-6">

                            <div class="radio"><label>
                                    <input type="radio"  value="1" name="status" checked> 启用</label></div>
                            <div class="radio"><label>
                                    <input type="radio"  value="0" name="status" > 禁用</label>
                            </div>

                        </div>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group">
                    <div class="col-sm-4 col-sm-offset-2">
                        <input type="hidden"  value="0" name="sn" id="sn">
                        <button class="btn btn-success btn-s-xs" type="submit">提交</button>

                    </div>
                </div>
                </form>
            </div>
        </section>
    </section>
</section>
<div class="modal fade tip-modal">
    <div class="modal-dialog modal-sm">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                        aria-hidden="true">&times;</span></button>

            </div>
            <div class="modal-body">
                <table id="table_show">
                    <?php
                    if($organization){
                        $oaService->dealwithData($organization);
                    }
                    ?>
                </table>
                    <a href="javascript:void(0)"  data-dismiss="modal" aria-label="Close" class="btn btn-dark right" style="margin-top: 10px;" >确定</a>
            </div>

        </div>
    </div>
</div>
<script>
    function show_user(){
        $('.tip-modal').modal();
    }
    function show_level(id) {
        $(".tr_" + id).toggle();
        $("tr[class='tr_" + id + "']").each(function () {
            if($(this).is(":hidden")){
                var sub_id = $(this).data('id');
                $(".tr_" + sub_id).hide();
                $("tr[class='tr_" + sub_id + "']").each(function () {
                    var sub_two_id = $(this).data('id');
                    $(".tr_" + sub_two_id).hide();

                })
            }

        })
    }
    function select_name(name,realname,sn) {
        $("#username").val(name);
        $("#nickname").val(realname);
        $("#sn").val(sn);

    }

</script>