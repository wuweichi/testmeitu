<?php
/**
 * Created by PhpStorm.
 * Date: 16-7-13
 * Time: 下午3:41
 */
$this->title = '编辑';
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 权限管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['user/index']);?>">用户列表</a></li>
            <li class="active"><a href="#">用户编辑</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">编辑</h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 用户编辑</header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo \Yii::$app->urlManager->createUrl(['user/detail']);?>" method="post" data-validate="parsley">
                <div class="form-group"><label class="col-lg-2 control-label">用户名：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="wangmoumou" data-required="true" class="form-control parsley-validated"  name="username" disabled value="<?php echo $model['username'];?>">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">姓名：</label>
                    <div class="col-sm-10">
                        <input type="text" data-required="true" placeholder="王某某" class="form-control parsley-validated" name="nickname" value="<?php echo $model['nickname'];?>">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">新密码：</label>
                    <div class="col-sm-10">
                        <input id="pwd" class="form-control parsley-validated" type="password" name="new_password" placeholder="minlength = 6" data-minlength="6">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">再次输入密码：</label>
                    <div class="col-sm-10">
                        <input class="form-control parsley-validated" type="password"  name="confirm_password" data-equalto="#pwd" placeholder="minlength = 6" data-minlength="6">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">角色：</label>

                    <div class="col-sm-10">
                        <select class="form-control parsley-validated" data-required="true" name="role_id">
                            <option value="">请选择</option>
                            <?php foreach ($role as $val): ?>
                                <option value="<?php echo $val['role_id'];?>" <?php echo $model['role_id'] == $val['role_id'] ? 'selected' : ''; ?>><?php echo $val['role_name'];?> </option>
                            <?php endforeach;?>
                        </select>
                    </div>
                </div>

                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">状态：</label>
                    <div class="col-sm-10">
                        <div class="col-sm-6"> <!-- radio -->

                            <div class="radio"><label>
                                    <input type="radio"  value="1" name="status" <?php echo $model['status'] == 1 ? 'checked' : ''; ?>> 启用</label></div>
                            <div class="radio"><label>
                                    <input type="radio"  value="0" name="status" <?php echo $model['status'] == 0 ? 'checked' : ''; ?>> 禁用</label>
                            </div>

                        </div>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group">
                    <div class="col-sm-4 col-sm-offset-2">
                        <input type="hidden" value="<?php echo $model['id'];?>" name="id">
                        <button class="btn btn-success btn-s-xs" type="submit">提交</button>

                    </div>
                </div>
                </form>
            </div>
        </section>
    </section>
</section>