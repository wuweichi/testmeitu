<?php
/**
 * Created by PhpStorm.
 * Date: 16-7-18
 * Time: 下午2:53
 */
$this->title = isset($model) ? "编辑" : "添加";
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 权限管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['menu/index']);?>">菜单列表</a></li>
            <li class="active"><a href="#">菜单<?php echo isset($model) ? "编辑" : "添加" ?></a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none"><?php echo isset($model) ? "编辑" : "添加" ?></h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 菜单<?php echo isset($model) ? "编辑" : "添加" ?></header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo isset($model) ? \Yii::$app->urlManager->createUrl(['menu/detail']) : \Yii::$app->urlManager->createUrl(['menu/add']) ?>" method="post" data-validate="parsley">
                <div class="form-group"><label class="col-lg-2 control-label">名称：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="添加用户" data-required="true" class="form-control parsley-validated"  name="name" data-minlength="3"  value="<?php echo isset($model) ? $model['name'] : ''?>">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">上一级id：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="0" data-required="true" class="form-control parsley-validated"  name="parent_id" data-minlength="1"  value="<?php echo isset($model) ? $model['parent_id'] : ''?>">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">地址：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="/menu/add" data-required="true" class="form-control parsley-validated"  name="url" data-minlength="3"  value="<?php echo isset($model) ? $model['url'] : ''?>">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">排序：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="0" data-required="true" class="form-control parsley-validated"  name="ordinal" data-minlength="1"  value="<?php echo isset($model) ? $model['ordinal'] : '0'?>">
                    </div>
                </div>

                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">状态：</label>
                    <div class="col-sm-10">
                        <div class="col-sm-6"> <!-- radio -->

                            <div class="radio">
                                <label><input type="radio"  value="1" name="is_show" <?php echo !isset($model) || $model['is_show']==1 ? 'checked':''?>> 显示</label>
                            </div>
                            <div class="radio">
                                <label> <input type="radio"  value="0" name="is_show" <?php echo isset($model) && $model['is_show']==0 ? 'checked':''?>> 不显示</label>
                            </div>



                        </div>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group">
                    <div class="col-sm-4 col-sm-offset-2">
                        <input type="hidden" name="id" value="<?php echo isset($model) ? $model['id'] : 0 ?>">
                        <button class="btn btn-success btn-s-xs" type="submit">提交</button>

                    </div>
                </div>
                </form>
            </div>
        </section>
    </section>
</section>