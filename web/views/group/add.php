<?php
/**
 * Created by PhpStorm.
 * Date: 16-7-8
 * Time: 下午2:37
 */
$this->title = '任务添加';
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 项目管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['group/list']);?>">项目列表</a></li>
            <li class="active"><a href="#">添加项目</a></li>
        </ul>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 添加项目</header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo \Yii::$app->urlManager->createUrl(['group/add']);?>" method="post" data-validate="parsley">
                <div class="form-group"><label class="col-lg-2 control-label">项目名称：</label>
                    <div class="col-sm-10">
                        <input value="<?php echo empty($groupDetail['name']) ? "" : $groupDetail['name'];?>" type="text" autocomplete="off"  data-required="true" class="form-control parsley-validated"  name="name">
                    </div>
                </div>
                
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">关联角色：</label>
                	<div class="col-sm-3">
                        <select class="form-control m-b" name="role_id[]" multiple="multiple">
                        	<?php if (! empty($roleList)):?>
                        		<?php foreach ($roleList as $_role):?>
                        		<option value="<?php echo $_role['role_id'];?>" <?php echo empty($relationRole[$_role['role_id']]) ? "" : "selected";?>> <?php echo $_role['role_name'];?></option>
                        		<?php endforeach;?>
                        	<?php endif;?>
                        </select>
                        <p class="form-control-static">按住<font color="red">ctrl</font>或者<font color="red">shift</font>多选</p>
                    </div>
                </div>
                
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">状态：</label>
                    <div class="col-sm-10">
                        <div class="col-sm-6"> <!-- radio -->
                            <?php foreach ($statusList as $_statusValue => $_statusName) :?>
                        		<div class="radio">
                        			<label>
                                        <input type="radio"  value="<?php echo $_statusValue;?>" name="status" <?php echo ( !empty($groupDetail['status']) && $_statusValue == $groupDetail['status']) ? "checked" : "";?>> <?php echo $_statusName;?>
                                    </label>
                                </div>
                        	<?php endforeach;?>
                        </div>
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