<?php
/**
 * Created by PhpStorm.
 * Date: 16-7-8
 * Time: 上午9:15
 */
$this->title = '计划任务服务器列表';
use yii\widgets\LinkPager;

?>

<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 服务器管理</a></li>
            <li><a href="#">服务器列表</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">服务器列表<a class="btn btn-s-md btn-primary" href="<?php echo \Yii::$app->urlManager->createUrl(['server/add']);?>" style="float: right;margin-right: 20px">添加服务器IP</a></h3></div>
        <section class="panel panel-default">
            <header class="panel-heading"> 服务器列表 </header>
            <div class="row text-sm wrapper">
                <form id="w0" class="form-horizontal form-inline" action="/index.php" method="get">
                    <input name="r" value="crontab/index" type="hidden">

                    <div class="form-group" >
                        <label class="col-sm-3 control-label">名称：</label>
                        <div class="col-sm-9">
                            <input type="text"  class="form-control" name="title" value="<?php echo $title;?>">
                        </div>
                    </div>

                    <button  class="btn btn-s-md btn-success"  type="submit">查询</button>
                </form>
            </div>
            <div class="table-responsive">
                <table class="table table-striped b-t b-light text-sm">
                    <thead>
                    <tr>
                        <th><input type="checkbox"> ID</th>
                        <th>服务器别名</th>
                        <th>执行服务器IP</th>
                        <th>状态</th>
                        <th>操作</th>
                        <th>操作人</th>
                    </tr>
                    </thead>
                    <tbody>
                    <?php foreach ($models as $model): ?>
                        <tr>
                            <td><input type="checkbox" name="post[]" value="<?php echo $model['id']; ?>" class="item"> <?php echo $model['id']; ?></td>
                            <td><?php echo $model['name']; ?></td>
                            <td><?php echo $model['ip']; ?></td>
                            <td><?php echo $model['status']; ?></td>
                            <td><?php echo $model['admin_id']; ?></td>
                            <td>
                                <a href="<?php echo \Yii::$app->urlManager->createUrl(['server/detail', 'id' => $model['id']]);?>">编辑</a>
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
                            <?php //echo LinkPager::widget(['pagination' => $pages]); ?>
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

