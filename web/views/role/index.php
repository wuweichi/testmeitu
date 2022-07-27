<?php
/**
 * Created by PhpStorm.
 */

$this->title = '角色列表';
use yii\widgets\LinkPager;

?>

<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 权限管理</a></li>
            <li><a href="#">角色列表</a></li>
        </ul>
        <div class="m-b-md">
            <h3 class="m-b-none">角色列表
                <a class="btn btn-s-md btn-primary" href="<?php echo \Yii::$app->urlManager->createUrl(['role/add']); ?>"
                   style="float: right;margin-right: 20px">添加角色</a></h3></div>
        <section class="panel panel-default">
            <header class="panel-heading"> 角色列表</header>
            <div class="row text-sm wrapper">
                <form id="w0" class="form-horizontal form-inline" action="/index.php" method="get">
                    <input name="r" value="role/index" type="hidden">

                <div class="form-group"  style="width: 25%">
                    <label class="col-sm-3 control-label">角色名称：</label>
                    <div class="col-sm-9">
                        <input type="text"  class="form-control" name="role_name" value="<?php echo $role_name;?>">
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
                        <th style="text-align: center">名称</th>
						<th style="text-align: center">上级管理名称</th>
                        <th style="text-align: center">操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    <?php foreach ($models as $_roleID => $model): ?>
                        <tr>
                            <td><input type="checkbox" name="post[]" value="<?php echo $model['role_id']; ?>" class="item"> <?php echo $model['role_id']; ?></td>
                            <td style="text-align: center"><?php echo $model['role_name']; ?></td>
                            <td style="text-align: center"><?php echo $roleNameList[$model['manage_role_id']];?></td>
                            <td style="text-align: center">
                                <a href="<?php echo \Yii::$app->urlManager->createUrl(['role/detail', 'role_id' => $model['role_id']]); ?>">编辑</a>
                                |
                                <a href="<?php echo \Yii::$app->urlManager->createUrl(['role/permission', 'role_id' => $model['role_id']]); ?>">分配权限</a>
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
            $.post("<?php echo \Yii::$app->urlManager->createUrl(['role/delete']); ?>", {ids: data[0]}, function (result) {
                show_alert(result.message,result.status,1);
            }, 'json');
        }
    });
</script>

