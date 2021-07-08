<?php
$this->title = '任务添加';
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 任务管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['crontab/index']);?>">任务列表</a></li>
            <li class="active"><a href="#">任务添加</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">添加</h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 任务添加</header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo \Yii::$app->urlManager->createUrl(['crontab/add']);?>" method="post" data-validate="parsley">
                <input type="hidden" name="id" value="<?php echo empty($taskDetail['id']) ? 0 : $taskDetail['id'];?>">
                <input type="hidden" value="<?php echo $taskDetail['target_ip'] ? $taskDetail['target_ip']: '';?>" name="old_target_ip">
                
                <div class="form-group"><label class="col-lg-2 control-label">标题：</label>
                    <div class="col-sm-10">
                        <input type="text" autocomplete="off"  data-required="true" class="form-control parsley-validated"  name="title" value="<?php echo empty($taskDetail['title']) ? '' : $taskDetail['title'];?>">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">简介：</label>
                    <div class="col-sm-10">
                        <textarea autocomplete="off" rows="6" class="form-control" name="intro"><?php echo empty($taskDetail['intro']) ? '' : $taskDetail['intro'];?></textarea>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">IP：</label>
                    <div class="col-sm-10">
                        <select class="form-control m-b parsley-validated"  data-required="true"  name="target_ip">
                        	<option value>请选择</option>
                        	<?php if (! empty($serverList)):?>
                        		<?php foreach ($serverList as $_server):?>
                        		<option value="<?php echo $_server['ip'];?>" <?php echo ($_server['ip'] == $taskDetail['target_ip']) ? "selected" : "";?>> <?php echo $_server['name'];?></option>
                        		<?php endforeach;?>
                        	<?php endif;?>
                        </select>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">命令：</label>
                    <div class="col-sm-10">
                        <input type="text" autocomplete="off" data-required="true" class="form-control parsley-validated" name="command" value="<?php echo empty($taskDetail['command']) ? "" : $taskDetail['command'];?>"></div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">执行时间：</label>
                    <div class="col-sm-10">
                        <input type="text" autocomplete="off" data-required="true" placeholder="例如：*/2 2秒，**/2 2分钟 "class="form-control parsley-validated" name="planning_time" value="<?php echo empty($taskDetail['planning_time']) ? "" : $taskDetail['planning_time'];?>">
                    </div>
                </div>
                
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">任务所属分组：</label>
                    <div class="col-sm-10">
                        <select class="form-control m-b parsley-validated" name="group_id" data-required="true">
                        	<option value>请选择</option>
                        	<?php if (! empty($groupNameList)):?>
                        		<?php foreach ($groupNameList as $_groupID => $_groupName):?>
                        		<option value="<?php echo $_groupID;?>" <?php echo ($_groupID == $taskDetail['group_id']) ? "selected" : "";?>> <?php echo $_groupName;?></option>
                        		<?php endforeach;?>
                        	<?php endif;?>
                        </select>
                    </div>
                </div>
                
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">状态：</label>
                    <div class="col-sm-10">
                        <div class="col-sm-6"> <!-- radio -->
                        	<?php foreach ($statusList as $_statusValue => $_statusName) :?>
                        		<div class="radio">
                        			<label>
                                        <input type="radio"  value="<?php echo $_statusValue;?>" name="status" <?php echo ($_statusValue == $taskDetail['status']) ? "checked" : "";?>> <?php echo $_statusName;?>
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