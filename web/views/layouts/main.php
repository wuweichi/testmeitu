<?php
/* @var $this \yii\web\View */
/* @var $content string */

use yii\helpers\Html;
use yii\widgets\Breadcrumbs;
use app\assets\AppAsset;

AppAsset::register($this);
?>
<?php $this->beginPage() ?>
<!DOCTYPE html>
<html class="app">
<head>
    <meta charset="<?= Yii::$app->charset ?>"/>
    <title><?= Html::encode($this->title) ?></title>
    <meta name="description"
          content="app, web app, responsive, admin dashboard, admin, flat, flat ui, ui kit, off screen nav"/>
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1"/>
    <?php $this->head() ?>

<script src="/front/js/jquery-1.7.0.min.js" cache="false"></script>
</head>
<body>

<?php $this->beginBody() ?>
<section class="vbox">
    <?= $this->render('top') ?>
    <section>
        <section class="hbox stretch">
            <?= $this->render('left') ?>
            <section id="content">
                <?= $content ?>
            </section>
        </section>
    </section>

</section>
<div class="modal fade tip-modal-alert" id="tip-modal-alert">
    <div class="modal-dialog modal-sm">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                        aria-hidden="true">&times;</span></button>
                <h4 class="modal-title" id="">提示</h4>
            </div>
            <div class="modal-body">
                <p class="mt20 f16 text-center " id="message_alert"></p>
            </div>
        </div>
    </div>
</div>
<?php $this->endBody() ?>
<script>
    function show_alert(message, status ='success',reload = '',time=2000) {
        $('.tip-modal-alert').modal();
        if (status == 'error') {
            $("#message_alert").html('<span class="alert-warning">' + message + '</span>');
        } else {
            $("#message_alert").html('<span class="alert-success">' + message + '</span>');
        }

        setTimeout(function () {
            $(".tip-modal-alert").modal("hide")
        }, time);
        if (reload) {
            window.location.reload();
        }
    }
    function select(){
        var ids = "";
        var n = 0;
        $('.item').each(function() {
            if(this.checked){
                ids += $(this).val() + ',';
                n += 1;
            }
        })
        return [ids.substring(0, ids.length - 1),n];
    }
</script>
</body>
</html>
<?php $this->endPage() ?>
