<?php

$this->title = '任务详细页';
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 任务管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['crontab/index', 'id' => $model['id']]);?>">任务列表</a></li>
            <li><a href="#">任务编辑</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">任务详情页</h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 详情页</header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo \Yii::$app->urlManager->createUrl(['crontab/detail']);?>" method="post" data-validate="parsley">
                <div class="form-group"><label class="col-lg-2 control-label">标题：</label>
                    <div class="col-sm-10">
                        <input type="text" data-required="true" class="form-control parsley-validated"  name="title" value="<?php echo $model['title'] ? $model['title']:'';?>">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">简介：</label>
                    <div class="col-sm-10">
                        <textarea placeholder="Type your message"  rows="6" class="form-control " name="intro"><?php echo $model['intro'] ? $model['intro']:'';?></textarea>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">IP：</label>
                    <div class="col-sm-10">
                        <input type="text"  data-required="true" class="form-control parsley-validated" value="<?php echo $model['target_ip'] ? $model['target_ip']:'';?>" name="target_ip">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">命令：</label>
                    <div class="col-sm-10">
                        <input type="text" data-required="true" class="form-control parsley-validated" value="<?php echo $model ? $model['command']:'';?>" name="command"></div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">执行时间：</label>
                    <div class="col-sm-10">
                        <input type="text" data-required="true" placeholder="例如：*/2 2秒，**/2 2分钟 "class="form-control parsley-validated" value="<?php echo $model['planning_time'] ? $model['planning_time']:'';?>" name="planning_time">
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">任务所属分组：</label>
                    <div class="col-sm-10">
                        <select class="form-control m-b" name="group_id">
                        	<?php if (! empty($groupList)):?>
                        		<?php foreach ($groupList as $_group):?>
                        		<option value="<?php echo $_group['id'];?>"> <?php echo $_group['name'];?></option>
                        		<?php endforeach;?>
                        	<?php endif;?>
                        </select>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">状态：</label>

                    <div class="col-sm-10">
                        <div class="col-sm-6"> <!-- radio -->
                            <div class="radio"><label>
                                    <input type="radio" value="0" name="status" <?php echo $model['status'] == 0 ? 'checked' : ''; ?>>
                                    未上线</label></div>
                            <div class="radio"><label>
                                    <input type="radio" value="1" name="status"  <?php echo $model['status'] == 1 ? 'checked' : ''; ?>>
                                    上线</label></div>
                            <div class="radio"><label>
                                    <input type="radio" value="2" name="status"  <?php echo $model['status'] == 2 ? 'checked' : ''; ?>>
                                    已删除</label></div>

                        </div>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group">
                    <div class="col-sm-4 col-sm-offset-2">
                        <input type="hidden" value="<?php echo $model['id'] ? $model['id']: '';?>" name="id">
                        <input type="hidden" value="<?php echo $model['target_ip'] ? $model['target_ip']: '';?>" name="old_target_ip">
                        <button class="btn btn-success btn-s-xs" type="submit">Submit</button>

                    </div>
                </div>
                </form>
            </div>
        </section>
    </section>
</section>