<?php
/**
 * Created by PhpStorm.
 * User: lq
 * Date: 16-7-14
 * Time: 下午3:19
 */
use yii\bootstrap\ActiveForm;
$this->title = isset($model) ? "编辑" : "添加";
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 权限管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['role/index']);?>">角色列表</a></li>
            <li class="active"><a href="#">角色<?php echo isset($model) ? "编辑" : "添加" ?></a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none"><?php echo isset($model) ? "编辑" : "添加" ?></h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 角色<?php echo isset($model) ? "编辑" : "添加" ?></header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo isset($model) ? \Yii::$app->urlManager->createUrl(['role/detail']) : \Yii::$app->urlManager->createUrl(['role/add']) ?>" method="post" data-validate="parsley">
                <div class="form-group"><label class="col-lg-2 control-label">名称：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="技术组" data-required="true" class="form-control parsley-validated"  name="role_name" data-minlength="3"  value="<?php echo isset($model) ? $model['role_name'] : ''?>">
                    </div>
                </div>

                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group">
                    <div class="col-sm-4 col-sm-offset-2">
                        <input type="hidden" name="role_id" value="<?php echo isset($model) ? $model['role_id'] : 0 ?>">
                    </div>
                </div>
                
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">关联管理组：</label>
                	<div class="col-sm-10">
                        <select class="form-control m-b" name="manage_role_id">
                        	<?php if (! empty($manageRoleList)):?>
                        		<?php foreach ($manageRoleList as $_role):?>
                        		<option value="<?php echo $_role['role_id'];?>"> <?php echo $_role['role_name'];?></option>
                        		<?php endforeach;?>
                        	<?php endif;?>
                        </select>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group">
                    <div class="col-sm-4 col-sm-offset-2">
                        <button class="btn btn-success btn-s-xs" type="submit">提交</button>

                    </div>
                </div>

                </form>
            </div>
        </section>
    </section>
</section>