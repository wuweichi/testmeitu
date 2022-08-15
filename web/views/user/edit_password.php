<?php
/**
 * Created by PhpStorm.
 * Date: 16-7-13
 * Time: 下午3:41
 */
$this->title = '修改密码';
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 权限管理</a></li>
            <li><a href="#">用户列表</a></li>
            <li class="active"><a href="#">修改密码</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">修改密码</h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 修改密码</header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo \Yii::$app->urlManager->createUrl(['user/edit_password']);?>" method="post" data-validate="parsley">


                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">新密码：</label>
                    <div class="col-sm-10">
                        <input id="pwd" class="form-control parsley-validated" type="password" data-required="true"  name="new_password" placeholder="minlength = 6" data-minlength="6">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">再次输入密码：</label>
                    <div class="col-sm-10">
                        <input class="form-control parsley-validated" type="password" data-required="true"  name="confirm_password" data-equalto="#pwd" placeholder="minlength = 6" data-minlength="6">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group">
                    <div class="col-sm-4 col-sm-offset-2">
                        <input type="hidden" value="<?php echo $id;?>" name="id">
                        <button class="btn btn-success btn-s-xs" type="submit">提交</button>

                    </div>
                </div>
                </form>
            </div>
        </section>
    </section>
</section>