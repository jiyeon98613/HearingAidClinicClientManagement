package model

type Patient struct {
    CID   string    #ex: c_1...
    Name string     #한글
    Age  int        
    Adress string   #한글
    Contact string  #ex: 010-0000-0000
    HearingAid_ID string    #ex: h_1...
    ReceiverTypeLR string # MM MS SM 등
    EarTipTypeLR string #ex: SM 등
}
