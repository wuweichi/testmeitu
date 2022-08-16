<?php
/**
 * Created by PhpStorm.
 * User: lq
 * Date: 16-7-8
 * Time: 上午9:15
 */

$this->title = '用户列表';
use yii\widgets\LinkPager;
?>

<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 权限管理</a></li>
            <li><a href="#">用户列表</a></li>
        </ul>
        <div class="m-b-md">
            <h3 class="m-b-none">用户列表
                <a class="btn btn-s-md btn-primary" href="<?php echo \Yii::$app->urlManager->createUrl(['user/add']); ?>"
                                                        style="float: right;margin-right: 20px">添加帐号</a></h3></div>
        <section class="panel panel-default">
            <header class="panel-heading"> 用户列表</header>
            <div class="row text-sm wrapper">
                <form id="w0" class="form-horizontal form-inline" action="/index.php" method="get">
                    <input name="r" value="user/index" type="hidden">

                    <div class="form-group" >
                        <label class="col-sm-3 control-label">帐号：</label>
                        <div class="col-sm-9">
                            <input type="text"  class="form-control" name="username" value="<?php echo $username;?>">
                        </div>
                    </div>
                    <div class="form-group" style="width:16%">
                        <label class="col-sm-3 control-label"> 角色：</label>
                        <div class="col-sm-9">
                            <select class="input-sm form-control input-s-sm inline" name="role_id">
                                <option value="">请选择</option>
                                <?php foreach ($role as $val): ?>
                                    <option value="<?php echo $val['role_id'];?>" <?php echo $role_id == $val['role_id'] ? 'selected' : ''; ?>><?php echo $val['role_name'];?> </option>
                                <?php endforeach;?>

                            </select>
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
                        <th>帐号</th>
                        <th>姓名</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    <?php foreach ($models as $model): ?>
                        <tr>
                            <td><input type="checkbox" name="post[]" value="<?php echo $model['id']; ?>" class="item"> <?php echo $model['id']; ?></td>
                            <td><?php echo $model['username']; ?></td>
                            <td><?php echo $model['nickname']; ?></td>
                            <td><a href="<?php echo \Yii::$app->urlManager->createUrl(['user/detail', 'id' => $model['id']]); ?>">编辑</a></td>
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
            $.post("<?php echo \Yii::$app->urlManager->createUrl(['user/delete']); ?>", {ids: data[0]}, function (result) {
                show_alert(result.message,result.status,1);
            }, 'json');
        }
    });
</script>

