<?php
/**
 * Created by PhpStorm.
 * Date: 16-7-8
 * Time: 上午9:15
 */

?>

<header class="bg-black dk header navbar navbar-fixed-top-xs">
    <div class="navbar-header aside-md">
        <a href="#" class="navbar-brand" data-toggle="fullscreen">
            <img src="/front/images/logo.png" class="m-r-sm">
            计划任务管理
        </a>

    </div>

    <ul class="nav navbar-nav navbar-right hidden-xs nav-user">

        <li class="dropdown"><a href="#" class="dropdown-toggle" data-toggle="dropdown"> <span
                    class="thumb-sm avatar pull-left"> <img src="<?php echo 'https://www.gravatar.com/avatar/'.md5(Yii::$app->user->identity->username).'?d=monsterid';?>"> </span> <?php echo Yii::$app->user->identity->nickname;?> <b
                    class="caret"></b> </a>
            <ul class="dropdown-menu animated fadeInRight"><span class="arrow top"></span>
                <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['user/edit_password']); ?>">修改密码</a></li>
                <li class="divider"></li>
                <li> <a href="<?php echo \Yii::$app->urlManager->createUrl(['site/logout']); ?>">退出</a></li>
            </ul>
        </li>
    </ul>
</header>
