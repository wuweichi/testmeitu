<?php
/**
 * Created by PhpStorm.
 * Date: 16-7-20
 * Time: 下午2:38
 */
$this->title = isset($model) ? "编辑" : "添加";
?>
<section class="vbox">
    <section class="scrollable padder">
        <ul class="breadcrumb no-border no-radius b-b b-light pull-in">
            <li><a href="#"><i class="fa fa-home"></i> 权限管理</a></li>
            <li><a href="<?php echo \Yii::$app->urlManager->createUrl(['role/index']);?>">角色列表</a></li>
            <li class="active"><a href="#">权限分配</a></li>
        </ul>
        <div class="m-b-md"><h3 class="m-b-none">权限分配</h3></div>
        <section class="panel panel-default">
            <header class="panel-heading font-bold"> 权限分配</header>
            <div class="panel-body">
                <form id="w0" class="form-horizontal" action="<?php echo \Yii::$app->urlManager->createUrl(['role/permission','role_id' => $role_id]);?>" method="post" data-validate="parsley">

                    <table class="table table-striped b-t b-light text-sm">
                        <tbody>

                        <?php foreach ($menu as $key=>$value): ?>
                            <tr>
                                <td style="text-indent:<?php echo $value['count']==1 ? $value['count'].'0px': $value['count'].'9px'?>;">
                                    <input type='checkbox' parent_id="<?php echo $value['parent_id']; ?>"
                                           level='<?php echo $value['count']; ?>' name="menu[]"
                                           value="<?php echo $value['id']; ?>" onclick="javascript:checknode(this);" <?php echo $value['checked'];?>>
                                    &nbsp&nbsp<?php echo $value['name']; ?>
                                </td>

                            </tr>
                        <?php endforeach; ?>
                        </tbody>
                    </table>

                    <div class="form-group">
                        <div class="col-sm-4 col-sm-offset-2">
                            <input type="hidden" name="role_id"  value="<?php echo $role_id;?>">
                            <button class="btn btn-success btn-s-xs" type="submit">提交</button>

                        </div>
                    </div>
                </form>

            </div>
        </section>
    </section>
</section>

<script>
    function checknode(obj)
    {
        var chk = $("input[type='checkbox']");
        var count = chk.length;
        var num = chk.index(obj);
        var level_top = chk.eq(num).attr('level');
        var level_bottom =  chk.eq(num).attr('level');
        for (var i = num; i>=0; i--)
        {
            var le = chk.eq(i).attr('level');
            if(eval(le) < eval(level_top))
            {
                chk.eq(i).attr("checked",true);
                var level_top = level_top-1;
            }
        }
        for (var j = num+1; j<count; j++)
        {
            var le = chk.eq(j).attr('level');
            if(chk.eq(num).attr("checked")=='checked' || chk.eq(num).is(":checked")) {
                if(eval(le) > eval(level_bottom)) chk.eq(j).attr("checked",true);
                else if(eval(le) == eval(level_bottom)) break;
            }
            else {
                if(eval(le) > eval(level_bottom)) chk.eq(j).attr("checked",false);
                else if(eval(le) == eval(level_bottom)) break;
            }
        }


    }
</script>