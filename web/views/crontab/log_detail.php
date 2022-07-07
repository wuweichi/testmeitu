<?php

$this->title = '日志';
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 日志管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['crontabl/log', 'id' => $model['id']]);?>">日志列表</a></li>
            <li><a href="#">日志详情页</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">日志详情页</h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 详情页</header>
            <div class="panel-body">
                <form class="form-horizontal" method="get">
                    <div class="form-group"><label class="col-lg-2 control-label">任务ID：</label>

                        <div class="col-lg-10"><p class="form-control-static"><?php echo $model['job_id'];?></p>
                        </div>
                    </div>
                    <div class="line line-dashed line-lg pull-in"></div>
                    <div class="form-group"><label class="col-lg-2 control-label">执行命令：</label>

                        <div class="col-lg-10"><p class="form-control-static"><?php echo $crontab['command'];?></p>
                        </div>
                    </div>

                    <div class="line line-dashed line-lg pull-in"></div>
                    <div class="form-group"><label class="col-lg-2 control-label">开始时间：</label>

                        <div class="col-lg-10"><p class="form-control-static"><?php echo $model['start_at'];?></p>
                        </div>
                    </div>
                    <div class="line line-dashed line-lg pull-in"></div>
                    <div class="form-group"><label class="col-lg-2 control-label">结束时间：</label>

                        <div class="col-lg-10"><p class="form-control-static"><?php echo $model['end_at'];?></p>
                        </div>
                    </div>
                    <div class="line line-dashed line-lg pull-in"></div>
                    <div class="form-group"><label class="col-lg-2 control-label">内容：</label>

                        <div class="col-lg-10"><p class="form-control-static"><?php echo $model['log_detail'];?></p>
                        </div>
                    </div>
                    <div class="line line-dashed line-lg pull-in"></div>
                    <div class="form-group">
                        <div class="col-sm-4 col-sm-offset-2">
                            <button type="submit" class="btn btn-primary" onclick="javascript :history.back(-1);'">返回</button>

                        </div>
                    </div>
                </form>
            </div>
        </section>
    </section>
</section>