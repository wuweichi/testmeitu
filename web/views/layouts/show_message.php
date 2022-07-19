<!DOCTYPE html>
<html class="bg-dark">
<head>
    <meta charset="utf-8"/>
    <title>提示信息</title>
    <meta name="description"
          content="app, web app, responsive, admin dashboard, admin, flat, flat ui, ui kit, off screen nav"/>
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1"/>
    <link rel="stylesheet" href="/front/css/app.v2.css" type="text/css"/>
    <link rel="stylesheet" href="/front/css/font.css" type="text/css" cache="false"/>
    <!--[if lt IE 9]>
    <script src="/front/js/ie/html5shiv.js" cache="false"></script>
    <script src="/front/js/ie/respond.min.js" cache="false"></script>
    <script src="/front/js/ie/excanvas.js" cache="false"></script> <![endif]--></head>
<body>
<section id="content" class="m-t-lg wrapper-md animated fadeInDown" style="margin-top: 40px">
    <div class="container aside-xxl">
        <section class="panel panel-default m-t-lg bg-white">
            <header class="panel-heading text-center"><strong>提示信息</strong></header>

            <div class="form-group" style="margin-top: 15px">
                <p class="text-muted text-center">
                    <?php if (isset($successMessage)): ?>
                        <span class="fa fa-check text-active"></span>
                        <span class="btn-lg text-success"><?php echo $successMessage; ?></span>
                    <?php elseif(isset($errorMessage)): ?>
                        <span class="fa fa-check text-active"></span>
                        <span class="btn-lg text-danger"><?php echo $errorMessage; ?></span>
                    <?php else: ?>
                        <span class="glyphicon glyphicon-ok-sign text-success"></span>
                        <span class="btn-lg text-success">操作成功！</span>
                    <?php endif; ?>
                </p>

            </div>
            <div class="form-group">
                <p class="text-muted text-center">该页将在<b id="wait"><?php echo $sec;?></b>秒后自动跳转!</p>
            </div>
            <div class="form-group">
                <p class="text-muted text-center">
                    <?php if(isset($gotoUrl)):?>
                        <a href="<?php echo $gotoUrl?>" id="href" class="btn btn-s-md btn-primary">立即跳转</a>
                    <?php else:?>
                        <a href="javascript:void(0)" onclick="history.go(-1)" class="btn btn-s-md btn-primary">返回上一页</a>
                    <?php endif;?>
                </p>
            </div>


        </section>
    </div>
</section>
<script src="/front/js/app.v2.js"></script>
<script type="text/javascript">

    (function(){
        var wait = document.getElementById('wait'),href = document.getElementById('href').href;
        totaltime=parseInt(wait.innerHTML);
        var interval = setInterval(function(){
            var time = --totaltime;
            wait.innerHTML=""+time;
            if(time === 0) {
                location.href = href;
                clearInterval(interval);
            };
        }, 1000);
    })();

</script>
</html>
<?php exit();?>
