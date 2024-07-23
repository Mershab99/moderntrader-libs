from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Side(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    BUY: _ClassVar[Side]
    SELL: _ClassVar[Side]
    PUT: _ClassVar[Side]
    CALL: _ClassVar[Side]

class SECFilingType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    Form4: _ClassVar[SECFilingType]
    Form13: _ClassVar[SECFilingType]

class Brokerage(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    INTERACTIVE_BROKERS: _ClassVar[Brokerage]
BUY: Side
SELL: Side
PUT: Side
CALL: Side
Form4: SECFilingType
Form13: SECFilingType
INTERACTIVE_BROKERS: Brokerage

class SECFilingEvent(_message.Message):
    __slots__ = ("id", "db_filing_id", "timestamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    DB_FILING_ID_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    db_filing_id: int
    timestamp: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[int] = ..., db_filing_id: _Optional[int] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class OrderSignal(_message.Message):
    __slots__ = ("id", "strategy_id", "ticker", "side", "timestamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    STRATEGY_ID_FIELD_NUMBER: _ClassVar[int]
    TICKER_FIELD_NUMBER: _ClassVar[int]
    SIDE_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    strategy_id: int
    ticker: str
    side: Side
    timestamp: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[int] = ..., strategy_id: _Optional[int] = ..., ticker: _Optional[str] = ..., side: _Optional[_Union[Side, str]] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class OrderPlacement(_message.Message):
    __slots__ = ("id", "strategy_id", "user_id", "broker_account_id", "ticker", "side", "quantity", "timestamp")
    ID_FIELD_NUMBER: _ClassVar[int]
    STRATEGY_ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    BROKER_ACCOUNT_ID_FIELD_NUMBER: _ClassVar[int]
    TICKER_FIELD_NUMBER: _ClassVar[int]
    SIDE_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    id: int
    strategy_id: int
    user_id: str
    broker_account_id: int
    ticker: str
    side: Side
    quantity: float
    timestamp: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[int] = ..., strategy_id: _Optional[int] = ..., user_id: _Optional[str] = ..., broker_account_id: _Optional[int] = ..., ticker: _Optional[str] = ..., side: _Optional[_Union[Side, str]] = ..., quantity: _Optional[float] = ..., timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
