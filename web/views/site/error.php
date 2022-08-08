<?php

/* @var $this yii\web\View */
/* @var $name string */
/* @var $message string */
/* @var $exception Exception */

use yii\helpers\Html;
$this->title = $name;
?>
<!DOCTYPE html>
<html lang="en" class="bg-dark">
<head>
    <meta charset="utf-8"/>
    <title>错误日志</title>
    <meta name="description"
          content="app, web app, responsive, admin dashboard, admin, flat, flat ui, ui kit, off screen nav"/>
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1"/>
    <link rel="stylesheet" href="/front/css/app.v2.css" type="text/css"/>
    <link rel="stylesheet" href="/front/css/font.css" type="text/css" cache="false"/>
    <!--[if lt IE 9]>
    <script src="/front/js/ie/html5shiv.js" cache="false"></script>
    <script src="/front/js/ie/respond.min.js" cache="false"></script>
    <script src="/front/js/ie/excanvas.js" cache="false"></script>
    <![endif]-->
</head>
<body>
<section id="content">
    <div class="row m-n">
        <div class="col-sm-4 col-sm-offset-4">
            <div class="text-center m-b-lg">
                <h1 class="h text-white animated fadeInDownBig"><?php echo $code; ?></h1>
            </div>
            <div class="list-group m-b-sm bg-white m-b-lg  m-t-xs" style="padding: 10px 0px">
                <div class="form-group">
                    <p class="text-muted text-center h3"><?= Html::encode($this->title) ?></p>
                </div>
                <div class="form-group">
                    <p class="text-muted text-center h3"><?= nl2br(Html::encode($message)) ?></p>
                </div>
            </div>
        </div>
    </div>
</section>
</body>
</html>

