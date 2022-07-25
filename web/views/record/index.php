<?php
/**
 * Created by PhpStorm.
 * User: lq
 * Date: 16-7-8
 * Time: 上午9:15
 */
use yii\helpers\Html;

$this->title = '日志';
use yii\widgets\LinkPager;

?>

<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 日志管理</a></li>
            <li><a href="#">日志列表</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">日志列表</h3></div>
        <section class="panel panel-default">
            <header class="panel-heading"> 日志列表</header>

            <div class="table-responsive">
                <table class="table table-striped b-t b-light text-sm">
                    <thead>
                    <tr>
                        <th><input type="checkbox"> ID</th>
                        <th>任务ID</th>
                        <th>开始时间</th>
                        <th>结束时间</th>
                        <th>操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    <?php foreach ($models as $model): ?>
                        <tr>
                            <td><input type="checkbox" name="post[]" value="<?php echo $model['id']; ?>"></td>
                            <td><?php echo $model['job_id']; ?></td>
                            <td><?php echo $model['start_at']; ?></td>
                            <td><?php echo $model['end_at']; ?></td>
                            <td><a href="<?php echo \Yii::$app->urlManager->createUrl(['record/detail', 'id' => $model['id']]);?>">查看</a> </td>
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

