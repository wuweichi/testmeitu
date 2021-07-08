<?php

// comment out the following two lines when deployed to production
$allowEnv = ['local','test','online'];
$currentEnv = isset($_SERVER['ENV_VAR_MY'])  &&  in_array($_SERVER['ENV_VAR_MY'],$allowEnv) ? $_SERVER['ENV_VAR_MY'] : 'online';

require(__DIR__ . '/vendor/autoload.php');
require(__DIR__ . '/vendor/yiisoft/yii2/Yii.php');

$config = require(__DIR__ . '/config/web.php');

(new yii\web\Application($config))->run();