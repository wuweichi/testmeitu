<?php

/* @var $this yii\web\View */
/* @var $form yii\bootstrap\ActiveForm */
/* @var $model app\models\LoginForm */

use yii\helpers\Html;
use yii\bootstrap\ActiveForm;
$this->title = 'Login';
$this->params['breadcrumbs'][] = $this->title;
?>
<!DOCTYPE html>
<html lang="en" class="bg-dark">
<head>
    <meta charset="utf-8"/>
    <title>脚本监控</title>
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

<section id="content" class="m-t-lg wrapper-md animated fadeInUp">
    <div class="container aside-xxl"><a class="navbar-brand block" href="index.html">Admin Manage</a>
        <section class="panel panel-default bg-white m-t-lg">
            <header class="panel-heading text-center"><strong>Sign in</strong></header>


            <?php $form = ActiveForm::begin([
                'id' => 'login-form',
                'options' => ['class' => 'panel-body wrapper-lg'],
                'fieldConfig' => [
                    'template' => "<div style='clear: both'>{label}\n{input}\n<div class=\"col-lg-8\">{error}</div></div>\n",
                    'labelOptions' => ['class' => 'control-label'],
                    'inputOptions' => ['class' => 'form-control input-lg'],
                ],
            ]); ?>

            <?= $form->field($model, 'username')->textInput(['autofocus' => true]) ?>

            <?= $form->field($model, 'password')->passwordInput() ?>

            <?= $form->field($model, 'rememberMe')->checkbox([
                'template' => "<div class=\"checkbox\" style=\"display: none\">{input} {label}</div>\n<div class=\"col-lg-8\" style=\"display: none\">{error}</div>",
            ]) ?>

            <div class="form-group" style="margin-top: 40px">

                <?= Html::submitButton('登录', ['class' => 'btn btn-primary', 'name' => 'login-button']) ?>
                <a href="<?php echo $url;?>" class="btn btn-danger" style="float: right">OA一键登录</a>

            </div>

            <?php ActiveForm::end(); ?>
        </section>
    </div>
</section>
<!-- footer -->
<footer id="footer">
    <div class="text-center padder">
        <p>
            <small>计划任务后台<br>&copy; 2016</small>
        </p>
    </div>
</footer>
<!-- / footer -->
</body>
</html>

