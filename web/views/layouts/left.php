<?php
$currentRedirect = Yii::$app->requestedRoute;
$tools = new app\models\lib\LibTool;
$menuList = $tools->roleMenu();
?>
<!-- .aside -->
<aside class="bg-black lter aside-md hidden-print" id="nav">
    <section class="vbox">
        <header class="header bg-primary lter text-center clearfix">
            <div class="btn-group">
                <button type="button" class="btn btn-sm btn-dark btn-icon" title="New project">
                    <i class="fa fa-plus"></i>
                </button>
                <div class="btn-group hidden-nav-xs">
                    <button type="button" class="btn btn-sm btn-primary dropdown-toggle"
                            data-toggle="dropdown"> 切换项目<span class="caret"></span></button>
                    <ul class="dropdown-menu text-left">
                        <li><a href="#">任务管理系统</a></li>

                    </ul>
                </div>
            </div>
        </header>
        <section class="w-f scrollable">
            <div class="slim-scroll" data-height="auto" data-disable-fade-out="true" data-distance="0"
                 data-size="5px" data-color="#333333"> <!-- nav -->
                <nav class="nav-primary hidden-xs">
                    <ul class="nav">

                    <?php foreach($menuList as $value):?>

                        <li class="<?php echo $value['class'];?>">
                            <a href="#layout" class="active">
                                <i class="fa fa-columns icon">
                                    <b class="bg-warning"></b>
                                </i>
                                <span class="pull-right">
                                    <i class="fa fa-angle-down text"></i>
                                    <i class="fa fa-angle-up text-active"></i>
                                </span>
                                <span><?php echo $value['name'];?></span>
                            </a>
                            <ul class="nav lt">
                         <?php foreach($value['subMenu'] as $val):?>
                                <li class="<?php echo $val['class'];?>">
                                    <a href="<?php echo \Yii::$app->urlManager->createUrl([$val['url']]); ?>" class="<?php echo $val['class'];?>">
                                        <i class="fa fa-angle-right"></i>
                                        <span><?php echo $val['name'];?></span>
                                    </a>
                                </li>
                            <?php endforeach;?>

                            </ul>
                        </li>
                        <?php endforeach;?>

                    </ul>
                </nav>
                <!-- / nav --> </div>
        </section>
        <footer class="footer lt hidden-xs b-t b-black">
            <div id="chat" class="dropup">
                <section class="dropdown-menu on aside-md m-l-n">
                    <section class="panel bg-white">
                        <header class="panel-heading b-b b-light">Active chats</header>
                        <div class="panel-body animated fadeInRight"><p class="text-sm">No active chats.</p>

                            <p><a href="#" class="btn btn-sm btn-default">Start a chat</a></p></div>
                    </section>
                </section>
            </div>
            <div id="invite" class="dropup">
                <section class="dropdown-menu on aside-md m-l-n">
                    <section class="panel bg-white">
                        <header class="panel-heading b-b b-light"> John <i
                                class="fa fa-circle text-success"></i></header>
                        <div class="panel-body animated fadeInRight"><p class="text-sm">No contacts in your
                                lists.</p>

                            <p><a href="#" class="btn btn-sm btn-facebook"><i
                                        class="fa fa-fw fa-facebook"></i> Invite from Facebook</a></p></div>
                    </section>
                </section>
            </div>
            <a href="#nav" data-toggle="class:nav-xs" class="pull-right btn btn-sm btn-black btn-icon">
                <i class="fa fa-angle-left text"></i>
                <i class="fa fa-angle-right text-active"></i> </a>


        </footer>
    </section>
</aside>
<!-- /.aside -->
