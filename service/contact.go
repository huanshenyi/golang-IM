package service

import (
	"errors"
	"hello/model"
	"time"
)

type ContactService struct {

}
// 友達追加
func (service *ContactService) AddFriend(userid, dstid int64 ) error{
	// 自分を追加する場合
	if userid==dstid{
		return errors.New("自分を追加できません")
	}
	// 既に追加したかどうかを判断
	tmp := model.Contact{}
	// 既に追加したかどうか調べる
	//
	DbEngin.Where("ownerid = ?",userid).
		And("dstobj = ?",dstid).
		And("cate = ?",model.CONCAT_CATE_USER).
		Get(&tmp)
	//もしデータが一個存在すれば
	//count()
	//追加したユーザーであることが判明
	if tmp.Id>0{
		return errors.New("該当ユーザー既に追加しました")
	}
	//Session 二つのデータ追加，片方失敗する場合、全部失敗する
	session := DbEngin.NewSession()
	session.Begin()
	// 自分の
	_,e2 := session.InsertOne(model.Contact{
		Ownerid:userid,
		Dstobj:dstid,
		Cate:model.CONCAT_CATE_USER,
		Createat:time.Now(),
	})
	// 相手の
	_,e3 := session.InsertOne(model.Contact{
		Ownerid:dstid,
		Dstobj:userid,
		Cate:model.CONCAT_CATE_USER,
		Createat:time.Now(),
	})
	// errorがなければ、DBにコミットする
	if  e2==nil && e3==nil{
		//提出
		session.Commit()
		return nil
	}else{
		//roll-back
		session.Rollback()
		if e2!=nil{
			return e2
		}else{
			return e3
		}
	}
}




func (service *ContactService) SearchComunity(userId int64) ([]model.Community){
	conconts := make([]model.Contact,0)
	comIds :=make([]int64,0)

	DbEngin.Where("ownerid = ? and cate = ?",userId,model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _,v := range conconts{
		comIds = append(comIds,v.Dstobj);
	}
	coms := make([]model.Community,0)
	if len(comIds)== 0{
		return coms
	}
	DbEngin.In("id",comIds).Find(&coms)
	return coms
}

//加群
func (service *ContactService) JoinCommunity(userId,comId int64) error{
	cot := model.Contact{
		Ownerid:userId,
		Dstobj:comId,
		Cate:model.CONCAT_CATE_COMUNITY,
	}
	DbEngin.Get(&cot)
	if(cot.Id==0){
		cot.Createat = time.Now()
		_,err := DbEngin.InsertOne(cot)
		return err
	}else{
		return nil
	}


}
//建群
func (service *ContactService) CreateCommunity(comm model.Community) (ret model.Community,err error){
	if len(comm.Name)==0{
		err = errors.New("缺少群名称")
		return ret,err
	}
	if comm.Ownerid==0{
		err = errors.New("请先登录")
		return ret,err
	}
	com := model.Community{
		Ownerid:comm.Ownerid,
	}
	num,err := DbEngin.Count(&com)

	if(num>5){
		err = errors.New("一个用户最多只能创见5个群")
		return com,err
	}else{
		comm.Createat=time.Now()
		session := DbEngin.NewSession()
		session.Begin()
		_,err = session.InsertOne(&comm)
		if err!=nil{
			session.Rollback();
			return com,err
		}
		_,err =session.InsertOne(
			model.Contact{
				Ownerid:comm.Ownerid,
				Dstobj:comm.Id,
				Cate:model.CONCAT_CATE_COMUNITY,
				Createat:time.Now(),
			})
		if err!=nil{
			session.Rollback();
		}else{
			session.Commit()
		}
		return com,err
	}
}

//查找好友
func (service *ContactService) SearchFriend(userId int64) ([]model.User){
	conconts := make([]model.Contact,0)
	objIds :=make([]int64,0)
	DbEngin.Where("ownerid = ? and cate = ?",userId,model.CONCAT_CATE_USER).Find(&conconts)
	for _,v := range conconts{
		objIds = append(objIds,v.Dstobj);
	}
	coms := make([]model.User,0)
	if len(objIds)== 0{
		return coms
	}
	DbEngin.In("id",objIds).Find(&coms)
	return coms
}