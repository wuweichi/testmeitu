<?php
/**
 * Created by PhpStorm.
 * Date: 16-7-8
 * Time: 上午9:15
 */
$this->title = '计划任务项目列表';
use yii\widgets\LinkPager;

?>

<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 项目管理</a></li>
            <li><a href="#">项目列表</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">项目列表<a class="btn btn-s-md btn-primary" href="<?php echo \Yii::$app->urlManager->createUrl(['group/add']);?>" style="float: right;margin-right: 20px">添加项目</a></h3></div>
        <section class="panel panel-default">
            <div class="table-responsive">
                <table class="table table-striped b-t b-light text-sm">
                    <thead>
                    <tr>
                        <th><input type="checkbox"> ID</th>
                        <th>分组名</th>
                        <th>授权角色组</th>
                        <th>状态</th>
                        <th>操作人</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    <?php foreach ($groupList as $_group): ?>
                        <tr>
                            <td><input type="checkbox" name="post[]" value="<?php echo $_group['id']; ?>" class="item"> <?php echo $_group['id']; ?></td>
                            <td><?php echo $_group['name']; ?></td>
                            <td><?php echo implode(' | ', $_group['role_name']); ?></td>
                            <td><?php echo $_group['status']; ?></td>
                            <td><?php echo $_group['admin_id']; ?></td>
                            <td>
                                <a href="<?php echo \Yii::$app->urlManager->createUrl(['group/add', 'id' => $_group['id']]);?>">编辑</a>
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
            $.post("<?php echo \Yii::$app->urlManager->createUrl(['group/delete']); ?>", {ids: data[0]}, function (result) {
                show_alert(result.message,result.status,1);
            }, 'json');
        }
    });
</script>

