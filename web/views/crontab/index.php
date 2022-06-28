<?php

$this->title = '任务列表';
use yii\widgets\LinkPager;

?>

<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 任务管理</a></li>
            <li><a href="#">任务列表</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">任务列表<a class="btn btn-s-md btn-primary" href="<?php echo \Yii::$app->urlManager->createUrl(['crontab/add']);?>" style="float: right;margin-right: 20px">添加任务</a></h3></div>
        <section class="panel panel-default"> 
        	<div class="panel-body">
            	<form class="form-inline" role="form" action="/index.php" method="get"> 
            		<div class="form-group"> 
            			<input type="hidden" name="r" value="crontab/index">
            			<select class="input-sm form-control input-s-sm inline" name="target_ip">
                            	<option>全部</option>
                            	<?php if (! empty($serverList)):?>
                            		<?php foreach ($serverList as $_server):?>
                            		<option value="<?php echo $_server['ip'];?>"> <?php echo $_server['name'];?></option>
                            		<?php endforeach;?>
                            	<?php endif;?>
                        </select> 
            		</div> 

            		<button class="btn btn-success" type="submit" >查询</button>
            	</form> 
        	</div> 

            <div class="table-responsive">
                <table class="table table-striped b-t b-light text-sm">
                    <thead>
                    <tr>
                        <th><input type="checkbox"> ID</th>
                        <th>名称</th>
                        <th>执行服务器IP</th>
                        <th>需要执行的命令</th>
                        <th>计划时间</th>
                        <th>状态</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    <?php foreach ($models as $model): ?>
                        <tr>
                            <td><input type="checkbox" name="post[]" value="<?php echo $model['id']; ?>" class="item"> <?php echo $model['id']; ?></td>
                            <td><?php echo $model['title']; ?></td>
                            <td><?php echo $model['target_ip']; ?></td>
                            <td><?php echo $model['command']; ?></td>
                            <td><?php echo $model['planning_time']; ?></td>
                            <td>
                                <?php if($model['status']==0):?>
                                未上线
                                <?php elseif($model['status']==1):?>
                                <span class="label label-success">上线</span>
                                <?php else:?>
                                    <span class="label label-warning">已删除</span>
                                <?php endif;?>
                            </td>
                            <td>
                                <a href="<?php echo \Yii::$app->urlManager->createUrl(['crontab/add', 'id' => $model['id']]);?>">编辑</a>
                                |
                                <a href="<?php echo \Yii::$app->urlManager->createUrl(['crontab/log', 'id' => $model['id']]);?>">查看日志</a>
                            </td>
                        </tr>
                    <?php endforeach; ?>
                    </tbody>
                </table>

            </div>
            <footer class="panel-footer">
                <div class="row">
                    <div class="col-sm-4 hidden-xs">
                        <button class="btn btn-dark" id="delete">批量删除</button>
                    </div>
                </div>
                <div class="row">
                    <div class="col-sm-6 text-right text-center-xs">
                        <ul class="pagination pagination-sm m-t-none m-b-none">
                            <?php echo LinkPager::widget(['pagination' => $pages]); ?>
                        </ul>
                    </div>
                </div>
            </footer>

        </section>
    </section>
</section>
<script>
    $('#delete').click(function(){
        var data = select();
        if(data[1] == 0){
            show_alert('您还未选择任何项!','error');
            return false;
        }else{
            $.post("<?php echo \Yii::$app->urlManager->createUrl(['crontab/delete']); ?>", {ids: data[0]}, function (result) {
                show_alert(result.message,result.status,1);
            }, 'json');
        }
    });
</script>

