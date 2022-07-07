<?php
/**
 * Created by PhpStorm.
 * 任务计划
 */
namespace app\controllers;

use app\models\db\MCrontabList;
use app\models\db\Server;
use app\models\db\Group;

use Yii;
use yii\yiidata\Pagination;
use app\models\form\CrontabForm;
use app\models\db\MExecutionLog;
use app\models\db\RoleTaskGroup;

class CrontabController extends BackendController
{
    /**
     * 列表
     * @return string
     */
    public function actionIndex()
    {
        // 获取服务器列表
        $serverList = Server::find()->where([])
        ->orderBy('id DESC')
        ->asArray()
        ->all();
        
        $title = empty($_GET['title']) ? '' : $_GET['title'];
        $serverIP = empty($_GET['target_ip']) ? 0 : trim($_GET['target_ip']);
        
        // 获取所属分组
        $groupIDList = [];
        $groupList = RoleTaskGroup::getUserOwnGroup(Yii::$app->user->identity->role_id);
        if (! empty($groupList)) {
            foreach ($groupList as $_group) {
                $groupIDList[] = $_group['group_id'];
            }
        }
        $whereParams = ['in', 'group_id', $groupIDList];
        $query = MCrontabList::find()->where($whereParams);
        if ($title) {
            $query->andWhere(['like', 'title', $title]);
        }
        if (! empty($serverIP)) {
            $query->andWhere(['=','target_ip', $serverIP]);
        }
        $countQuery = clone $query;
        $pages = new Pagination(['totalCount' => $countQuery->count()]);
        $models = $query->offset($pages->offset)
            ->limit($pages->limit)
            ->orderBy('id DESC')->asArray()->all();

        return $this->render('index', [
            'models' => $models,
            'pages' => $pages,
            'title' => $title,
            'serverList' => $serverList
        ]);
    }

    /**
     * 详细页
     * @return string|\yii\web\Response
     */

    public function actionDetail()
    {

        $modelForm = new CrontabForm();
        //更新信息
        if (Yii::$app->request->post()) {
            $result = $modelForm->update(Yii::$app->request->post());
            if ($result) {
                return $this->success(['crontab/index'], '编辑 成功', 1);
            } else {
                return $this->error(['crontab/detail', 'id' => $_POST['id']], '修改失败', 1);
            }

        }
        $id = intval($_GET['id']);
        $model = MCrontabList::find()->where(['id' => $id])->asArray()->one();
        // 获取任务分组列表
        $groupList = Group::getUserOwnGroup(Yii::$app->user->identity->role_id);
        $viewParams['model'] = $model;
        $viewParams['id'] = $id;
        $viewParams['groupList'] = $groupList;
        return $this->render('detail', $viewParams);

    }

    public function actionAdd()
    {
        $id = empty($_GET['id']) ? 0 : intval($_GET['id']);
        $taskDetail = MCrontabList::find()->where(['id' => $id])->asArray()->one();
        // 获取服务器列表
        $serverList = Server::find()->where([])
            ->orderBy('id DESC')
            ->asArray()
            ->all();
        
        // 获取任务分组列表
        $groupNameList = [];
        $groupList = RoleTaskGroup::getUserOwnGroup(Yii::$app->user->identity->role_id);
        if (! empty($groupList)) {
            foreach ($groupList as $_group) {
                $_groupList[$_group['group_id']] = $_group['group_id'];
            }
            $groupList = Group::getGroupList($_groupList);
            foreach ($groupList as $_group) {
                $groupNameList[$_group['id']] = $_group['name'];
            }
        }
        //更新信息
        $modelForm = new CrontabForm();
        if (Yii::$app->request->post()) {
            if (! empty(Yii::$app->request->post()['id'])) {
                $modelForm->update(Yii::$app->request->post());
            } else {
                $modelForm->insert(Yii::$app->request->post());
            }
            
            return $this->success(['crontab/index'],'添加成功',1);
        }
        // 状态列表
        $statusList = $this->getAllStatus();
        $viewParams['serverList'] = $serverList;
        $viewParams['groupNameList'] = $groupNameList;
        $viewParams['taskDetail'] = $taskDetail;
        $viewParams['statusList'] = $statusList;
        return $this->render('add', $viewParams);

    }


    /**
     * 任务日志列表
     * @return string
     */
    public function actionLog()
    {
        if (!isset($_GET['id'])) {
            $where = '1 = 1';
        } else {
            $jobId = (int)$_GET['id'];
            $where = ['job_id' => $jobId];
        }
        $query = MExecutionLog::find()->where($where);
        $countQuery = clone $query;
        $pages = new Pagination(['totalCount' => $countQuery->count()]);
        $models = $query->offset($pages->offset)
            ->limit($pages->limit)
            ->orderBy('id DESC')->asArray()->all();
        return $this->render('log_index', [
            'models' => $models,
            'pages' => $pages,
            'job_id' => $jobId
        ]);
    }

    /**
     * 任务详情页
     * @return string
     */
    public function actionLog_detail()
    {
        $id = intval($_GET['id']);
        $model = MExecutionLog::find()->where(['id' => $id])->asArray()->one();
        $crontabModel = MCrontabList::find()->where(['id' => $model['job_id']])->asArray()->one();
        return $this->render('log_detail', [
            'model' => $model,
            'id' => $id,
            'crontab' => $crontabModel
        ]);

    }

    /**
     * 状态码
     * @param $status
     * @return mixed
     */
    public function getStatus($status)
    {
        $statusMessage = array('0' => '未上线', '1' => '上线', '2' => '已删除');
        return $statusMessage[$status];
    }
    
    /**
     * 获取所有支持的状态
     * @return string[]
     */
    public function getAllStatus()
    {
        $statusMessage = array('0' => '未上线', '1' => '上线', '2' => '已删除');
        return $statusMessage;
    }

    /**
     * 删除
     */
    public function actionDelete()
    {
        $ids = explode(',', $_POST['ids']);
        $result = MCrontabList::deleteAll(['id' => $ids]);
        if ($result) {
            MExecutionLog::deleteAll(['job_id' => $ids]);
            echo json_encode(array('message' => '删除成功', 'status' => 'success'));
        } else {
            echo json_encode(array('message' => '删除失败', 'status' => 'error'));
        }
        die();

    }

}