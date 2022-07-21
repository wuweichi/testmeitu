<?php
/**
 * Created by PhpStorm.
 * User: lq
 * Date: 16-7-18
 * Time: 上午10:13
 */
$this->title = '菜单列表';
use yii\widgets\LinkPager;

?>

<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 权限管理</a></li>
            <li><a href="#">菜单列表</a></li>
        </ul>
        <div class="m-b-md">
            <h3 class="m-b-none">菜单列表
                <a class="btn btn-s-md btn-primary" href="<?php echo \Yii::$app->urlManager->createUrl(['menu/add']); ?>"
                   style="float: right;margin-right: 20px">添加菜单</a></h3></div>
        <section class="panel panel-default">
            <header class="panel-heading"> 菜单列表</header>
            <div class="row text-sm wrapper">
                <form id="w0" class="form-horizontal form-inline" action="/index.php" method="get">
                    <input name="r" value="menu/index" type="hidden">

                <div class="form-group"  style="width: 25%">
                    <label class="col-sm-3 control-label">菜单名称：</label>
                    <div class="col-sm-9">
                        <input type="text"  class="form-control" name="name" value="<?php echo $name;?>">
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
                        <th>名称</th>
                        <th>父级ID</th>
                        <th>状态</th>
                        <th>地址</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    <?php foreach ($models as $model): ?>
                        <tr>
                            <td><input type="hidden" name="post[]" value="<?php echo $model['id']; ?>" class="item"> <?php echo $model['id']; ?></td>
                            <td><?php echo $model['name']; ?></td>
                            <td><?php echo $model['parent_id']; ?></td>
                            <td><?php echo $model['is_show'] == 1? '显示':'不显示'; ?></td>
                            <td><?php echo $model['url']; ?></td>
                            <td><a href="<?php echo \Yii::$app->urlManager->createUrl(['menu/detail', 'id' => $model['id']]); ?>">编辑</a> |
                                <a href="javascript:void(0)" onclick="delete_menu(<?php echo $model['id'];?>)">删除</a>
                             </td>
                        </tr>
                    <?php endforeach; ?>
                    </tbody>
                </table>

            </div>
            <footer class="panel-footer">

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
    function delete_menu(id){
        $.post("<?php echo \Yii::$app->urlManager->createUrl(['menu/delete']); ?>", {id:id}, function (result) {
            show_alert(result.message,result.status,1);
        }, 'json');

    }

</script>