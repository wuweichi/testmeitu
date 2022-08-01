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
            <li><a href="#"><i class="fa fa-home"></i> 任务管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['server/list']);?>">服务器列表</a></li>
            <li class="active"><a href="#">添加服务器</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">添加</h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 添加服务器</header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo \Yii::$app->urlManager->createUrl(['server/add']);?>" method="post" data-validate="parsley">
                <div class="form-group"><label class="col-lg-2 control-label">服务器名称：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="127.0.0.1" data-required="true" class="form-control parsley-validated"  name="title">
                    </div>
                </div>
                
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">IP：</label>
                    <div class="col-sm-10">
                        <input type="text" placeholder="127.0.0.1" data-required="true" class="form-control parsley-validated"  name="target_ip">
                    </div>
                </div>
                
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group"><label class="col-lg-2 control-label">状态：</label>
                    <div class="col-sm-10">
                        <div class="col-sm-6"> <!-- radio -->

                                <div class="radio"><label>
                                        <input type="radio"  value="0" name="status" checked> 未上线</label></div>
                                <div class="radio"><label>
                                        <input type="radio"  value="1" name="status" > 上线</label></div>
                                <div class="radio"><label>
                                        <input type="radio"  value="2" name="status" > 已删除</label></div>
                        </div>
                    </div>
                </div>
                <div class="line line-dashed line-lg pull-in"></div>
                <div class="form-group">
                    <div class="col-sm-4 col-sm-offset-2">
                        <button class="btn btn-success btn-s-xs" type="submit">Submit</button>

                    </div>
                </div>
                </form>
            </div>
        </section>
    </section>
</section>