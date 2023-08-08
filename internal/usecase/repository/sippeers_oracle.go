package repository

import (
	"context"
	"database/sql"
	"fmt"
	"myapp/internal/model"

	go_oracle "gitlabnew.nextcontact.ru/r.alfimov/go-oracle"
)

func NewSippeersOracleDB(ora *go_oracle.Oracle) *SippeersOracle {
	return &SippeersOracle{ora}
}

type SippeersOracle struct {
	*go_oracle.Oracle
}

func (r *SippeersOracle) InsertSippeerses(ctx context.Context, SippersTest []model.Sippers) error {

	sqlText := `INSERT INTO sippeers_test (NAME,            IPADDR                ,PORT                  ,REGSECONDS            ,DEFAULTUSER           ,FULLCONTACT           ,REGSERVER             ,USERAGENT             ,LASTMS                ,HOST                  ,TYPE                  ,CONTEXT               ,PERMIT                ,DENY                  ,SECRET                ,MD5SECRET             ,REMOTESECRET          ,TRANSPORT             ,DTMFMODE              ,DIRECTMEDIA           ,NAT                   ,CALLGROUP             ,PICKUPGROUP           ,LANGUAGE              ,DISALLOW              ,ALLOW                 ,INSECURE              ,TRUSTRPID             ,PROGRESSINBAND        ,PROMISCREDIR          ,USECLIENTCODE         ,ACCOUNTCODE           ,SETVAR                ,CALLERID              ,AMAFLAGS              ,CALLCOUNTER           ,BUSYLEVEL             ,ALLOWOVERLAP          ,ALLOWSUBSCRIBE        ,VIDEOSUPPORT          ,MAXCALLBITRATE        ,RFC2833COMPENSATE     ,MAILBOX               ,"SESSION-TIMERS"        ,"SESSION-EXPIRES"       ,"SESSION-MINSE"         ,"SESSION-REFRESHER"     ,T38PT_USERTPSOURCE    ,REGEXTEN              ,FROMDOMAIN            ,FROMUSER              ,QUALIFY               ,DEFAULTIP             ,RTPTIMEOUT            ,RTPHOLDTIMEOUT        ,SENDRPID              ,OUTBOUNDPROXY         ,CALLBACKEXTENSION     ,TIMERT1               ,TIMERB                ,QUALIFYFREQ           ,CONSTANTSSRC          ,CONTACTPERMIT         ,CONTACTDENY           ,USEREQPHONE           ,TEXTSUPPORT           ,FAXDETECT             ,BUGGYMWI              ,AUTH                  ,FULLNAME              ,TRUNKNAME             ,CID_NUMBER            ,CALLINGPRES           ,MOHINTERPRET          ,MOHSUGGEST            ,PARKINGLOT            ,HASVOICEMAIL          ,SUBSCRIBEMWI          ,VMEXTEN               ,AUTOFRAMING           ,RTPKEEPALIVE          ,CALL_LIMIT            ,G726NONSTANDARD       ,IGNORESDPVERSION      ,ALLOWTRANSFER         ,DYNAMIC               ,PATH                  ,SUPPORTPATH           ,CANREINVITE           ,AVPF                  ,FORCE_AVP ,          ICESUPPORT,          RTCP_MUX,            ENCRYPTION ,         DTLSENABLE ,         DTLSVERIFY           ,DTLSCERTFILE          ,DTLSPRIVATEKEY        ,DTLSSETUP             ,DTLSREKEY)
  VALUES         (:1,            :2                ,:3                  ,:4            ,:5           ,:6           ,:7             ,:8             ,:9                ,:10                  ,:11                  ,:12               ,:13                ,:14                  ,:15                ,:16             ,:17          ,:18             ,:19              ,:20           ,:21                   ,:22             ,:23           ,:24              ,:25              ,:26                 ,:27              ,:28             ,:29        ,:30          ,:31         ,:32           ,:33                ,:34              ,:35              ,:36           ,:37             ,:38          ,:39        ,:40          ,:41        ,:42     ,:43               ,:44        ,:45       ,:46         ,:47     ,:48    ,:49              ,:50            ,:51              ,:52               ,:53             ,:54            ,:55        ,:56              ,:57         ,:58     ,:59               ,:60                ,:61           ,:62          ,:63         ,:64           ,:65           ,:66           ,:67             ,:68              ,:69                  ,:70              ,:71             ,:72            ,:73           ,:74          ,:75            ,:76            ,:77          ,:78          ,:79               ,:80           ,:81          ,:82            ,:83       ,:84      ,:85         ,:86               ,:87                  ,:88           ,:89           ,:90                  ,:91,          :92,          :93,            :94 ,        :95 ,         :96           ,:97          ,:98        ,:99             ,:100)`

	values := make([]interface{}, len(SippersTest))
	for idx, SippersTest := range SippersTest {
		values[idx] = []interface{}{SippersTest.Name, SippersTest.IpAddr, SippersTest.Port, SippersTest.RegSeconds, SippersTest.DefaultUser, SippersTest.FullContact, SippersTest.RegServer, SippersTest.UserAgent, SippersTest.LastMS, SippersTest.Host, SippersTest.Type, SippersTest.Context, SippersTest.Permit, SippersTest.Deny, SippersTest.Secret, SippersTest.MD5Secret, SippersTest.RemoteSecret, SippersTest.Transport, SippersTest.DtmFMode, SippersTest.DirectMedia, SippersTest.Nat, SippersTest.CallGroup, SippersTest.PickUpGroup, SippersTest.Language, SippersTest.Disallow, SippersTest.Allow, SippersTest.Insecure, SippersTest.Trustrprid, SippersTest.ProgressInBand, SippersTest.PromisCredir, SippersTest.UseClientCode, SippersTest.AccountCode, SippersTest.Setvar, SippersTest.CallerId, SippersTest.AmaFlags, SippersTest.CallCounter, SippersTest.BusyLevel, SippersTest.AllowOverLap, SippersTest.AllowSubscribe, SippersTest.VidoSupport, SippersTest.MaxCallBitrate, SippersTest.RFC2833Compensate, SippersTest.MailBox, SippersTest.SessionTimers, SippersTest.SessionExpires, SippersTest.SessionMinse, SippersTest.SessionRefresher, SippersTest.T38prUserTPSourse, SippersTest.RegExten, SippersTest.FromDomain, SippersTest.FromUser, SippersTest.Qualify, SippersTest.DefaultTip, SippersTest.RtpTimeout, SippersTest.RtpHoldTimeout, SippersTest.SendRpid, SippersTest.OutBoundProxy, SippersTest.CallBackExtension, SippersTest.Timert1, SippersTest.TimerB, SippersTest.QualifyFreq, SippersTest.ConstantSsrc, SippersTest.ContactPermit, SippersTest.ContactDeny, SippersTest.UseReqPhone, SippersTest.TextSupport, SippersTest.FaxDetect, SippersTest.BuggyMwi, SippersTest.Auth, SippersTest.FullName, SippersTest.TrunkName, SippersTest.CIDNumber, SippersTest.CallIngpres, SippersTest.MohinTerpet, SippersTest.MohSuggest, SippersTest.ParkingLot, SippersTest.HasVoiceMail, SippersTest.SubscribeMWI, SippersTest.Vmexten, SippersTest.AutoFraming, SippersTest.RtpKeepAlive, SippersTest.CallLimit, SippersTest.G726NonStandard, SippersTest.IgnoreSDPVersion, SippersTest.AllowTransfer, SippersTest.Dynamic, SippersTest.Path, SippersTest.SupportPath, SippersTest.CanReinvite, SippersTest.Avpf, SippersTest.ForceAvp, SippersTest.IceSupport, SippersTest.RtcpMux, SippersTest.Encryption, SippersTest.DtlsEnable, SippersTest.DtlsVerify, SippersTest.DtlsCertfile, SippersTest.DtlsPrivateKey, SippersTest.DtlsSetup, SippersTest.DtlsRekey}
	}
	if err := go_oracle.InsertMany(ctx, r.Oracle, sqlText, values); err != nil {
		return fmt.Errorf("Oracle - InsertSippeerses - oracle.InsertMany: %w", err)
	}

	return nil
}

func (r *SippeersOracle) InsertSippeers(ctx context.Context, SippersTest model.Sippers) error {

	sqlText := `INSERT INTO sippeers_test (NAME,            IPADDR                ,PORT                  ,REGSECONDS            ,DEFAULTUSER           ,FULLCONTACT           ,REGSERVER             ,USERAGENT             ,LASTMS                ,HOST                  ,TYPE                  ,CONTEXT               ,PERMIT                ,DENY                  ,SECRET                ,MD5SECRET             ,REMOTESECRET          ,TRANSPORT             ,DTMFMODE              ,DIRECTMEDIA           ,NAT                   ,CALLGROUP             ,PICKUPGROUP           ,LANGUAGE              ,DISALLOW              ,ALLOW                 ,INSECURE              ,TRUSTRPID             ,PROGRESSINBAND        ,PROMISCREDIR          ,USECLIENTCODE         ,ACCOUNTCODE           ,SETVAR                ,CALLERID              ,AMAFLAGS              ,CALLCOUNTER           ,BUSYLEVEL             ,ALLOWOVERLAP          ,ALLOWSUBSCRIBE        ,VIDEOSUPPORT          ,MAXCALLBITRATE        ,RFC2833COMPENSATE     ,MAILBOX               ,"SESSION-TIMERS"        ,"SESSION-EXPIRES"       ,"SESSION-MINSE"         ,"SESSION-REFRESHER"     ,T38PT_USERTPSOURCE    ,REGEXTEN              ,FROMDOMAIN            ,FROMUSER              ,QUALIFY               ,DEFAULTIP             ,RTPTIMEOUT            ,RTPHOLDTIMEOUT        ,SENDRPID              ,OUTBOUNDPROXY         ,CALLBACKEXTENSION     ,TIMERT1               ,TIMERB                ,QUALIFYFREQ           ,CONSTANTSSRC          ,CONTACTPERMIT         ,CONTACTDENY           ,USEREQPHONE           ,TEXTSUPPORT           ,FAXDETECT             ,BUGGYMWI              ,AUTH                  ,FULLNAME              ,TRUNKNAME             ,CID_NUMBER            ,CALLINGPRES           ,MOHINTERPRET          ,MOHSUGGEST            ,PARKINGLOT            ,HASVOICEMAIL          ,SUBSCRIBEMWI          ,VMEXTEN               ,AUTOFRAMING           ,RTPKEEPALIVE          ,CALL_LIMIT            ,G726NONSTANDARD       ,IGNORESDPVERSION      ,ALLOWTRANSFER         ,DYNAMIC               ,PATH                  ,SUPPORTPATH           ,CANREINVITE           ,AVPF                  ,FORCE_AVP ,          ICESUPPORT,          RTCP_MUX,            ENCRYPTION ,         DTLSENABLE ,         DTLSVERIFY           ,DTLSCERTFILE          ,DTLSPRIVATEKEY        ,DTLSSETUP             ,DTLSREKEY)
  VALUES         (:1,            :2                ,:3                  ,:4            ,:5           ,:6           ,:7             ,:8             ,:9                ,:10                  ,:11                  ,:12               ,:13                ,:14                  ,:15                ,:16             ,:17          ,:18             ,:19              ,:20           ,:21                   ,:22             ,:23           ,:24              ,:25              ,:26                 ,:27              ,:28             ,:29        ,:30          ,:31         ,:32           ,:33                ,:34              ,:35              ,:36           ,:37             ,:38          ,:39        ,:40          ,:41        ,:42     ,:43               ,:44        ,:45       ,:46         ,:47     ,:48    ,:49              ,:50            ,:51              ,:52               ,:53             ,:54            ,:55        ,:56              ,:57         ,:58     ,:59               ,:60                ,:61           ,:62          ,:63         ,:64           ,:65           ,:66           ,:67             ,:68              ,:69                  ,:70              ,:71             ,:72            ,:73           ,:74          ,:75            ,:76            ,:77          ,:78          ,:79               ,:80           ,:81          ,:82            ,:83       ,:84      ,:85         ,:86               ,:87                  ,:88           ,:89           ,:90                  ,:91,          :92,          :93,            :94 ,        :95 ,         :96           ,:97          ,:98        ,:99             ,:100)`

	values := []interface{}{SippersTest.Name, SippersTest.IpAddr, SippersTest.Port, SippersTest.RegSeconds, SippersTest.DefaultUser, SippersTest.FullContact, SippersTest.RegServer, SippersTest.UserAgent, SippersTest.LastMS, SippersTest.Host, SippersTest.Type, SippersTest.Context, SippersTest.Permit, SippersTest.Deny, SippersTest.Secret, SippersTest.MD5Secret, SippersTest.RemoteSecret, SippersTest.Transport, SippersTest.DtmFMode, SippersTest.DirectMedia, SippersTest.Nat, SippersTest.CallGroup, SippersTest.PickUpGroup, SippersTest.Language, SippersTest.Disallow, SippersTest.Allow, SippersTest.Insecure, SippersTest.Trustrprid, SippersTest.ProgressInBand, SippersTest.PromisCredir, SippersTest.UseClientCode, SippersTest.AccountCode, SippersTest.Setvar, SippersTest.CallerId, SippersTest.AmaFlags, SippersTest.CallCounter, SippersTest.BusyLevel, SippersTest.AllowOverLap, SippersTest.AllowSubscribe, SippersTest.VidoSupport, SippersTest.MaxCallBitrate, SippersTest.RFC2833Compensate, SippersTest.MailBox, SippersTest.SessionTimers, SippersTest.SessionExpires, SippersTest.SessionMinse, SippersTest.SessionRefresher, SippersTest.T38prUserTPSourse, SippersTest.RegExten, SippersTest.FromDomain, SippersTest.FromUser, SippersTest.Qualify, SippersTest.DefaultTip, SippersTest.RtpTimeout, SippersTest.RtpHoldTimeout, SippersTest.SendRpid, SippersTest.OutBoundProxy, SippersTest.CallBackExtension, SippersTest.Timert1, SippersTest.TimerB, SippersTest.QualifyFreq, SippersTest.ConstantSsrc, SippersTest.ContactPermit, SippersTest.ContactDeny, SippersTest.UseReqPhone, SippersTest.TextSupport, SippersTest.FaxDetect, SippersTest.BuggyMwi, SippersTest.Auth, SippersTest.FullName, SippersTest.TrunkName, SippersTest.CIDNumber, SippersTest.CallIngpres, SippersTest.MohinTerpet, SippersTest.MohSuggest, SippersTest.ParkingLot, SippersTest.HasVoiceMail, SippersTest.SubscribeMWI, SippersTest.Vmexten, SippersTest.AutoFraming, SippersTest.RtpKeepAlive, SippersTest.CallLimit, SippersTest.G726NonStandard, SippersTest.IgnoreSDPVersion, SippersTest.AllowTransfer, SippersTest.Dynamic, SippersTest.Path, SippersTest.SupportPath, SippersTest.CanReinvite, SippersTest.Avpf, SippersTest.ForceAvp, SippersTest.IceSupport, SippersTest.RtcpMux, SippersTest.Encryption, SippersTest.DtlsEnable, SippersTest.DtlsVerify, SippersTest.DtlsCertfile, SippersTest.DtlsPrivateKey, SippersTest.DtlsSetup, SippersTest.DtlsRekey}

	if err := go_oracle.Insert(ctx, r.Oracle, sqlText, values); err != nil {
		return fmt.Errorf("Oracle - InsertSippeers - oracle.Insert: %w", err)
	}

	return nil
}

func (r *SippeersOracle) ListSippeersLogin(ctx context.Context) ([]model.ListSippersLogin, error) {

	sqlText := `SELECT name FROM sippeers_test`

	sippersLogin, err := go_oracle.SelectMany[model.ListSippersLogin](ctx, r.Oracle, sqlText, []interface{}{})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Oracle - ListSippeersLogin - go_oracle.SelectMany: %w", err)
	}
	return sippersLogin, nil
}
