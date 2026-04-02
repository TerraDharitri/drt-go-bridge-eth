(module
  (type (;0;) (func (param i32 i32)))
  (type (;1;) (func (result i32)))
  (type (;2;) (func (param i32 i32) (result i32)))
  (type (;3;) (func))
  (type (;4;) (func (param i32) (result i32)))
  (type (;5;) (func (param i32)))
  (type (;6;) (func (param i32 i32 i32) (result i32)))
  (type (;7;) (func (param i32) (result i64)))
  (type (;8;) (func (param i64 i32 i32 i32 i32 i32) (result i32)))
  (type (;9;) (func (param i32 i32 i32 i32 i32 i32 i32 i32 i64 i64 i32) (result i32)))
  (type (;10;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;11;) (func (param i32 i64)))
  (type (;12;) (func (result i64)))
  (type (;13;) (func (param i64)))
  (type (;14;) (func (param i32 i32 i32)))
  (type (;15;) (func (param i32 i32 i32 i32)))
  (type (;16;) (func (param i32 i32 i32 i32 i32)))
  (type (;17;) (func (param i64 i32 i32 i32 i32)))
  (type (;18;) (func (param i32 i32) (result i64)))
  (type (;19;) (func (param i64 i32)))
  (type (;20;) (func (param i32 i32 i32 i64)))
  (type (;21;) (func (param i32 i64 i32 i32)))
  (import "env" "signalError" (func (;0;) (type 0)))
  (import "env" "mBufferNew" (func (;1;) (type 1)))
  (import "env" "mBufferAppend" (func (;2;) (type 2)))
  (import "env" "mBufferEq" (func (;3;) (type 2)))
  (import "env" "cleanReturnData" (func (;4;) (type 3)))
  (import "env" "managedGetBackTransfers" (func (;5;) (type 0)))
  (import "env" "mBufferGetLength" (func (;6;) (type 4)))
  (import "env" "managedCaller" (func (;7;) (type 5)))
  (import "env" "mBufferGetBytes" (func (;8;) (type 2)))
  (import "env" "isSmartContract" (func (;9;) (type 4)))
  (import "env" "managedOwnerAddress" (func (;10;) (type 5)))
  (import "env" "mBufferGetArgument" (func (;11;) (type 2)))
  (import "env" "mBufferAppendBytes" (func (;12;) (type 6)))
  (import "env" "managedSignalError" (func (;13;) (type 5)))
  (import "env" "smallIntGetUnsignedArgument" (func (;14;) (type 7)))
  (import "env" "getNumArguments" (func (;15;) (type 1)))
  (import "env" "mBufferFinish" (func (;16;) (type 4)))
  (import "env" "managedExecuteOnDestContext" (func (;17;) (type 8)))
  (import "env" "managedCreateAsyncCall" (func (;18;) (type 9)))
  (import "env" "mBufferGetByteSlice" (func (;19;) (type 10)))
  (import "env" "mBufferSetBytes" (func (;20;) (type 6)))
  (import "env" "bigIntSign" (func (;21;) (type 4)))
  (import "env" "mBufferCopyByteSlice" (func (;22;) (type 10)))
  (import "env" "mBufferFromBigIntUnsigned" (func (;23;) (type 2)))
  (import "env" "mBufferToBigIntUnsigned" (func (;24;) (type 2)))
  (import "env" "bigIntSetInt64" (func (;25;) (type 11)))
  (import "env" "mBufferStorageLoad" (func (;26;) (type 2)))
  (import "env" "mBufferStorageStore" (func (;27;) (type 2)))
  (import "env" "getGasLeft" (func (;28;) (type 12)))
  (import "env" "managedSCAddress" (func (;29;) (type 5)))
  (import "env" "checkNoPayment" (func (;30;) (type 3)))
  (import "env" "managedGetMultiDCDTCallValue" (func (;31;) (type 5)))
  (import "env" "getBlockRound" (func (;32;) (type 12)))
  (import "env" "smallIntFinishUnsigned" (func (;33;) (type 13)))
  (import "env" "managedWriteLog" (func (;34;) (type 0)))
  (import "env" "smallIntFinishSigned" (func (;35;) (type 13)))
  (import "env" "managedGetCallbackClosure" (func (;36;) (type 5)))
  (func (;37;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call 0
    unreachable)
  (func (;38;) (type 4) (param i32) (result i32)
    (local i32)
    call 1
    local.tee 1
    local.get 0
    call 2
    drop
    local.get 1)
  (func (;39;) (type 2) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 3
    i32.const 0
    i32.gt_s)
  (func (;40;) (type 5) (param i32)
    (local i32 i32 i32 i32 i32 i64 i32 i64)
    global.get 0
    i32.const 208
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    i32.const 64
    i32.add
    local.get 0
    call 41
    call 42
    call 43
    local.set 2
    local.get 1
    i32.const 112
    i32.add
    local.get 0
    call 44
    call 45
    local.get 1
    call 46
    call 43
    i32.store offset=132
    local.get 1
    i32.const 112
    i32.add
    i32.const 8
    i32.add
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.load offset=80
          local.get 1
          i32.load offset=120
          call 39
          br_if 0 (;@3;)
          local.get 1
          i32.const 56
          i32.add
          i32.const 131328
          i32.const 11
          call 47
          local.get 1
          i32.const 48
          i32.add
          local.get 1
          i32.load offset=56
          local.get 1
          i32.load offset=60
          local.get 1
          i32.const 64
          i32.add
          i32.const 16
          i32.add
          call 48
          local.get 1
          i32.const 136
          i32.add
          i32.const 16
          i32.add
          local.tee 4
          local.get 3
          i32.store
          local.get 1
          local.get 1
          i32.const 124
          i32.add
          i32.store offset=156
          local.get 1
          local.get 1
          i64.load offset=112
          i64.store offset=144
          local.get 1
          local.get 1
          i32.load offset=52
          local.tee 3
          i32.store offset=140
          local.get 1
          local.get 1
          i32.load offset=48
          local.tee 5
          i32.store offset=136
          local.get 1
          local.get 1
          i32.const 132
          i32.add
          i32.store offset=160
          local.get 1
          call 49
          i64.store offset=96
          local.get 1
          i32.const 168
          i32.add
          i32.const 16
          i32.add
          local.get 4
          i64.load
          i64.store
          local.get 1
          local.get 1
          i64.load offset=144
          local.tee 6
          i64.store offset=176
          local.get 1
          local.get 1
          i32.const 164
          i32.add
          local.tee 4
          i32.store offset=200
          local.get 1
          local.get 4
          i32.store offset=196
          local.get 1
          local.get 3
          i32.store offset=172
          local.get 1
          local.get 5
          i32.store offset=168
          local.get 1
          local.get 1
          i32.const 96
          i32.add
          i32.store offset=192
          block  ;; label = @4
            block  ;; label = @5
              local.get 6
              i64.const 0
              i64.ne
              br_if 0 (;@5;)
              local.get 1
              i32.const 32
              i32.add
              local.get 5
              local.get 3
              local.get 1
              i32.load offset=184
              local.get 1
              i32.load offset=188
              call 50
              local.get 1
              i32.load offset=36
              local.set 3
              local.get 1
              i32.load offset=32
              local.set 5
              call 51
              local.set 4
              local.get 1
              i64.load offset=96
              local.get 1
              i32.load offset=132
              local.get 4
              local.get 5
              local.get 3
              call 52
              br 1 (;@4;)
            end
            local.get 1
            i32.const 40
            i32.add
            local.get 5
            local.get 3
            local.get 1
            i32.load offset=132
            local.get 1
            i32.const 168
            i32.add
            i32.const 8
            i32.add
            call 53
            local.get 1
            i32.load offset=44
            local.set 3
            local.get 1
            i32.load offset=40
            local.set 5
            call 54
            local.set 4
            call 51
            local.set 7
            local.get 1
            i64.load offset=96
            local.get 4
            local.get 7
            local.get 5
            local.get 3
            call 52
          end
          call 4
          call 55
          local.tee 3
          call 55
          local.tee 5
          call 5
          local.get 5
          call 56
          i32.eqz
          br_if 2 (;@1;)
          block  ;; label = @4
            local.get 3
            call 57
            i32.const 1
            i32.ne
            br_if 0 (;@4;)
            local.get 1
            i32.const 96
            i32.add
            local.get 3
            call 58
            br 2 (;@2;)
          end
          i32.const 131172
          i32.const 30
          call 59
          unreachable
        end
        local.get 1
        i32.const 96
        i32.add
        i32.const 8
        i32.add
        local.get 3
        i64.load
        i64.store
        local.get 1
        local.get 1
        i64.load offset=112
        i64.store offset=96
      end
      local.get 0
      call 60
      local.tee 0
      local.get 0
      call 61
      local.set 6
      local.get 1
      i64.load offset=64
      local.set 8
      local.get 1
      i32.load offset=76
      local.set 4
      local.get 1
      i32.load offset=72
      local.set 5
      local.get 1
      i32.const 24
      i32.add
      i32.const 131104
      i32.const 17
      call 47
      local.get 1
      i32.load offset=24
      local.set 3
      local.get 1
      i32.load offset=28
      local.set 0
      call 62
      call 38
      local.tee 7
      local.get 5
      call 2
      drop
      local.get 0
      local.get 7
      call 63
      call 62
      call 38
      local.tee 5
      local.get 4
      call 2
      drop
      local.get 6
      local.get 5
      call 64
      local.get 8
      local.get 5
      call 64
      local.get 0
      local.get 5
      call 63
      local.get 1
      i32.const 136
      i32.add
      i32.const 16
      i32.add
      local.tee 5
      local.get 1
      i32.const 96
      i32.add
      i32.const 8
      i32.add
      i32.store
      local.get 1
      local.get 2
      i32.store offset=160
      local.get 1
      local.get 1
      i32.const 108
      i32.add
      i32.store offset=156
      local.get 1
      local.get 0
      i32.store offset=140
      local.get 1
      local.get 3
      i32.store offset=136
      local.get 1
      local.get 1
      i64.load offset=96
      i64.store offset=144
      local.get 1
      call 49
      i64.store offset=112
      local.get 1
      i32.const 168
      i32.add
      i32.const 16
      i32.add
      local.get 5
      i64.load
      i64.store
      local.get 1
      local.get 1
      i64.load offset=144
      local.tee 6
      i64.store offset=176
      local.get 1
      local.get 1
      i32.const 164
      i32.add
      local.tee 5
      i32.store offset=200
      local.get 1
      local.get 5
      i32.store offset=196
      local.get 1
      local.get 0
      i32.store offset=172
      local.get 1
      local.get 3
      i32.store offset=168
      local.get 1
      local.get 1
      i32.const 112
      i32.add
      i32.store offset=192
      block  ;; label = @2
        block  ;; label = @3
          local.get 6
          i64.const 0
          i64.ne
          br_if 0 (;@3;)
          local.get 1
          i32.const 8
          i32.add
          local.get 3
          local.get 0
          local.get 1
          i32.load offset=184
          local.get 1
          i32.load offset=188
          call 50
          local.get 1
          i32.load offset=12
          local.set 0
          local.get 1
          i32.load offset=8
          local.set 3
          call 51
          local.set 5
          local.get 1
          i64.load offset=112
          local.get 2
          local.get 5
          local.get 3
          local.get 0
          call 52
          br 1 (;@2;)
        end
        local.get 1
        i32.const 16
        i32.add
        local.get 3
        local.get 0
        local.get 2
        local.get 1
        i32.const 168
        i32.add
        i32.const 8
        i32.add
        call 53
        local.get 1
        i32.load offset=20
        local.set 0
        local.get 1
        i32.load offset=16
        local.set 3
        call 54
        local.set 5
        call 51
        local.set 2
        local.get 1
        i64.load offset=112
        local.get 5
        local.get 2
        local.get 3
        local.get 0
        call 52
      end
      call 4
      local.get 1
      i32.const 208
      i32.add
      global.set 0
      return
    end
    i32.const 131172
    i32.const 30
    call 59
    unreachable)
  (func (;41;) (type 0) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 52
    i32.add
    call 66
    local.get 2
    i32.const 8
    i32.add
    local.get 2
    i32.load offset=52
    local.get 2
    i32.load offset=60
    local.get 1
    call 80
    block  ;; label = @1
      local.get 2
      i64.load offset=8
      i64.const 0
      i64.ne
      br_if 0 (;@1;)
      i32.const 131202
      i32.const 13
      call 59
      unreachable
    end
    local.get 0
    local.get 2
    i64.load offset=16
    i64.store
    local.get 0
    i32.const 24
    i32.add
    local.get 2
    i32.const 40
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 16
    i32.add
    local.get 2
    i32.const 8
    i32.add
    i32.const 24
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 8
    i32.add
    local.get 2
    i32.const 8
    i32.add
    i32.const 16
    i32.add
    i64.load
    i64.store
    local.get 2
    i32.const 64
    i32.add
    global.set 0)
  (func (;42;) (type 1) (result i32)
    i32.const 131852
    i32.const 23
    call 97)
  (func (;43;) (type 4) (param i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      call 86
      local.tee 1
      call 6
      i32.const 32
      i32.eq
      br_if 0 (;@1;)
      local.get 0
      i32.const 131718
      i32.const 16
      call 126
      unreachable
    end
    local.get 1)
  (func (;44;) (type 4) (param i32) (result i32)
    (local i32)
    i32.const 131902
    i32.const 8
    call 97
    local.tee 1
    local.get 0
    call 135
    local.get 1)
  (func (;45;) (type 0) (param i32 i32)
    (local i32 i32 i64 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 12
    i32.add
    local.get 1
    call 85
    local.get 2
    i32.const 12
    i32.add
    local.get 1
    call 88
    local.set 3
    local.get 2
    i32.const 12
    i32.add
    local.get 1
    call 139
    local.set 4
    local.get 2
    i32.const 12
    i32.add
    local.get 1
    call 87
    local.set 5
    block  ;; label = @1
      local.get 2
      i32.load offset=16
      local.get 2
      i32.load offset=12
      i32.ne
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 2
        i32.load8_u offset=28
        i32.eqz
        br_if 0 (;@2;)
        i32.const 0
        i32.const 0
        i32.store offset=142016
        i32.const 0
        i32.const 0
        i32.store8 offset=142020
      end
      local.get 0
      local.get 5
      i32.store offset=12
      local.get 0
      local.get 3
      i32.store offset=8
      local.get 0
      local.get 4
      i64.store
      local.get 2
      i32.const 32
      i32.add
      global.set 0
      return
    end
    local.get 1
    i32.const 131072
    i32.const 14
    call 126
    unreachable)
  (func (;46;) (type 1) (result i32)
    i32.const 131875
    i32.const 27
    call 97)
  (func (;47;) (type 14) (param i32 i32 i32)
    local.get 1
    local.get 2
    call 97
    local.set 2
    local.get 0
    call 62
    i32.store offset=4
    local.get 0
    local.get 2
    i32.store)
  (func (;48;) (type 15) (param i32 i32 i32 i32)
    call 62
    drop
    local.get 2
    local.get 3
    i32.load
    call 38
    call 63
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;49;) (type 12) (result i64)
    (local i64)
    call 28
    local.tee 0
    i64.const -100000
    i64.add
    local.get 0
    local.get 0
    i64.const 100000
    i64.gt_u
    select)
  (func (;50;) (type 16) (param i32 i32 i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 5
    global.set 0
    local.get 5
    i32.const 24
    i32.add
    i32.const 131598
    i32.const 12
    call 47
    local.get 5
    i32.const 16
    i32.add
    local.get 5
    i32.load offset=24
    local.get 5
    i32.load offset=28
    local.get 3
    call 48
    local.get 5
    i32.const 8
    i32.add
    local.get 5
    i32.load offset=16
    local.get 5
    i32.load offset=20
    local.get 4
    call 116
    local.get 5
    local.get 5
    i32.load offset=8
    local.get 5
    i32.load offset=12
    local.get 1
    local.get 2
    call 117
    local.get 5
    i32.load offset=4
    local.set 3
    local.get 0
    local.get 5
    i32.load
    i32.store
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 5
    i32.const 32
    i32.add
    global.set 0)
  (func (;51;) (type 1) (result i32)
    i32.const -10
    i64.const 0
    call 25
    i32.const -10)
  (func (;52;) (type 17) (param i64 i32 i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    local.get 4
    call 55
    call 17
    drop)
  (func (;53;) (type 16) (param i32 i32 i32 i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 5
    global.set 0
    local.get 5
    i32.const 40
    i32.add
    i32.const 131583
    i32.const 15
    call 47
    local.get 5
    i32.const 32
    i32.add
    local.get 5
    i32.load offset=40
    local.get 5
    i32.load offset=44
    local.get 4
    i32.load offset=8
    call 48
    local.get 5
    i32.const 24
    i32.add
    local.get 5
    i32.load offset=32
    local.get 5
    i32.load offset=36
    local.get 4
    i64.load
    call 115
    local.get 5
    i32.const 16
    i32.add
    local.get 5
    i32.load offset=24
    local.get 5
    i32.load offset=28
    local.get 4
    i32.load offset=12
    call 116
    local.get 5
    i32.load offset=16
    local.set 6
    local.get 5
    i32.load offset=20
    local.set 4
    call 62
    drop
    local.get 4
    local.get 3
    call 38
    call 63
    local.get 5
    i32.const 8
    i32.add
    local.get 6
    local.get 4
    local.get 1
    local.get 2
    call 117
    local.get 5
    i32.load offset=12
    local.set 4
    local.get 0
    local.get 5
    i32.load offset=8
    i32.store
    local.get 0
    local.get 4
    i32.store offset=4
    local.get 5
    i32.const 48
    i32.add
    global.set 0)
  (func (;54;) (type 1) (result i32)
    (local i32)
    call 55
    local.tee 0
    call 29
    local.get 0)
  (func (;55;) (type 1) (result i32)
    (local i32)
    i32.const 0
    i32.const 0
    i32.load offset=132012
    i32.const -1
    i32.add
    local.tee 0
    i32.store offset=132012
    local.get 0)
  (func (;56;) (type 4) (param i32) (result i32)
    local.get 0
    call 21
    i32.eqz)
  (func (;57;) (type 4) (param i32) (result i32)
    local.get 0
    call 120
    i32.const 4
    i32.shr_u)
  (func (;58;) (type 0) (param i32 i32)
    (local i32 i32 i64 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 16
    i32.add
    i64.const 0
    i64.store
    local.get 2
    i64.const 0
    i64.store offset=8
    local.get 1
    i32.const 0
    local.get 2
    i32.const 8
    i32.add
    i32.const 16
    call 101
    local.set 1
    local.get 2
    i32.const 0
    i32.store offset=28
    local.get 2
    i32.const 8
    i32.add
    local.get 2
    i32.const 28
    i32.add
    call 129
    local.set 3
    local.get 2
    i32.const 8
    i32.add
    local.get 2
    i32.const 28
    i32.add
    call 130
    local.set 4
    local.get 2
    i32.const 8
    i32.add
    local.get 2
    i32.const 28
    i32.add
    call 129
    local.set 5
    block  ;; label = @1
      local.get 1
      i32.eqz
      br_if 0 (;@1;)
      i32.const 131625
      i32.const 29
      call 0
      unreachable
    end
    local.get 0
    local.get 5
    i32.store offset=12
    local.get 0
    local.get 3
    i32.store offset=8
    local.get 0
    local.get 4
    i64.store
    local.get 2
    i32.const 32
    i32.add
    global.set 0)
  (func (;59;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call 37
    unreachable)
  (func (;60;) (type 4) (param i32) (result i32)
    (local i32)
    i32.const 131121
    i32.const 8
    call 97
    local.tee 1
    local.get 0
    call 135
    local.get 1)
  (func (;61;) (type 18) (param i32 i32) (result i64)
    (local i32 i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i64.const 0
    i64.store offset=8
    block  ;; label = @1
      local.get 0
      call 86
      local.tee 3
      call 6
      local.tee 0
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 1
      i32.const 131072
      i32.const 14
      call 126
      unreachable
    end
    local.get 3
    i32.const 0
    local.get 2
    i32.const 8
    i32.add
    local.get 0
    i32.sub
    i32.const 8
    i32.add
    local.get 0
    call 101
    drop
    local.get 2
    i64.load offset=8
    local.set 4
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 4
    i64.const 56
    i64.shl
    local.get 4
    i64.const 65280
    i64.and
    i64.const 40
    i64.shl
    i64.or
    local.get 4
    i64.const 16711680
    i64.and
    i64.const 24
    i64.shl
    local.get 4
    i64.const 4278190080
    i64.and
    i64.const 8
    i64.shl
    i64.or
    i64.or
    local.get 4
    i64.const 8
    i64.shr_u
    i64.const 4278190080
    i64.and
    local.get 4
    i64.const 24
    i64.shr_u
    i64.const 16711680
    i64.and
    i64.or
    local.get 4
    i64.const 40
    i64.shr_u
    i64.const 65280
    i64.and
    local.get 4
    i64.const 56
    i64.shr_u
    i64.or
    i64.or
    i64.or)
  (func (;62;) (type 1) (result i32)
    (local i32)
    call 55
    local.tee 0
    i32.const 131960
    i32.const 0
    call 20
    drop
    local.get 0)
  (func (;63;) (type 0) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.const 24
    i32.shl
    local.get 1
    i32.const 65280
    i32.and
    i32.const 8
    i32.shl
    i32.or
    local.get 1
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 1
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    i32.store offset=12
    local.get 0
    local.get 2
    i32.const 12
    i32.add
    i32.const 4
    call 12
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;64;) (type 19) (param i64 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    i64.const 56
    i64.shl
    local.get 0
    i64.const 65280
    i64.and
    i64.const 40
    i64.shl
    i64.or
    local.get 0
    i64.const 16711680
    i64.and
    i64.const 24
    i64.shl
    local.get 0
    i64.const 4278190080
    i64.and
    i64.const 8
    i64.shl
    i64.or
    i64.or
    local.get 0
    i64.const 8
    i64.shr_u
    i64.const 4278190080
    i64.and
    local.get 0
    i64.const 24
    i64.shr_u
    i64.const 16711680
    i64.and
    i64.or
    local.get 0
    i64.const 40
    i64.shr_u
    i64.const 65280
    i64.and
    local.get 0
    i64.const 56
    i64.shr_u
    i64.or
    i64.or
    i64.or
    i64.store offset=8
    local.get 1
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    call 127
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;65;) (type 5) (param i32)
    (local i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 80
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    i32.const 36
    i32.add
    call 66
    block  ;; label = @1
      local.get 1
      i32.load offset=36
      local.tee 2
      local.get 0
      call 67
      local.tee 3
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 24
      i32.add
      local.get 1
      i32.load offset=40
      local.tee 4
      local.get 3
      call 68
      local.get 1
      i32.load offset=28
      local.set 5
      local.get 1
      i32.load offset=24
      local.set 6
      local.get 1
      i32.const 48
      i32.add
      local.get 4
      call 69
      block  ;; label = @2
        block  ;; label = @3
          local.get 6
          i32.eqz
          br_if 0 (;@3;)
          local.get 1
          i32.const 16
          i32.add
          local.get 4
          local.get 6
          call 68
          local.get 4
          local.get 6
          local.get 1
          i32.load offset=16
          local.get 5
          call 70
          br 1 (;@2;)
        end
        local.get 1
        local.get 5
        i32.store offset=52
      end
      block  ;; label = @2
        block  ;; label = @3
          local.get 5
          i32.eqz
          br_if 0 (;@3;)
          local.get 1
          i32.const 8
          i32.add
          local.get 4
          local.get 5
          call 68
          local.get 4
          local.get 5
          local.get 6
          local.get 1
          i32.load offset=12
          call 70
          br 1 (;@2;)
        end
        local.get 1
        local.get 6
        i32.store offset=56
      end
      local.get 4
      i32.const 131669
      i32.const 11
      local.get 3
      call 71
      call 72
      local.get 4
      local.get 3
      call 73
      drop
      local.get 4
      i32.const 131680
      i32.const 6
      local.get 3
      call 71
      call 72
      local.get 1
      local.get 1
      i32.load offset=48
      i32.const -1
      i32.add
      i32.store offset=48
      local.get 4
      local.get 1
      i32.const 48
      i32.add
      call 74
      local.get 2
      local.get 0
      call 75
      call 72
      local.get 1
      i32.const 48
      i32.add
      local.get 1
      i32.load offset=44
      local.tee 4
      local.get 0
      call 76
      local.get 4
      local.get 0
      call 77
      call 78
    end
    local.get 0
    call 79
    call 78
    local.get 1
    i32.const 80
    i32.add
    global.set 0)
  (func (;66;) (type 5) (param i32)
    (local i32 i32 i32)
    i32.const 131812
    i32.const 20
    call 97
    local.tee 1
    call 38
    local.set 2
    local.get 1
    call 38
    local.set 3
    local.get 0
    local.get 1
    i32.store offset=4
    local.get 0
    local.get 3
    i32.store
    local.get 0
    local.get 2
    i32.store offset=8)
  (func (;67;) (type 2) (param i32 i32) (result i32)
    (local i64)
    block  ;; label = @1
      local.get 0
      local.get 1
      call 75
      local.tee 1
      local.get 1
      call 61
      local.tee 2
      i64.const 4294967296
      i64.lt_u
      br_if 0 (;@1;)
      local.get 1
      i32.const 131072
      i32.const 14
      call 126
      unreachable
    end
    local.get 2
    i32.wrap_i64)
  (func (;68;) (type 14) (param i32 i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    i32.const 12
    i32.add
    local.get 1
    i32.const 131669
    i32.const 11
    local.get 2
    call 71
    local.tee 2
    call 85
    local.get 3
    i32.const 12
    i32.add
    local.get 2
    call 123
    local.set 1
    local.get 3
    i32.const 12
    i32.add
    local.get 2
    call 123
    local.set 4
    block  ;; label = @1
      local.get 3
      i32.load offset=16
      local.get 3
      i32.load offset=12
      i32.ne
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 3
        i32.load8_u offset=28
        i32.eqz
        br_if 0 (;@2;)
        i32.const 0
        i32.const 0
        i32.store offset=142016
        i32.const 0
        i32.const 0
        i32.store8 offset=142020
      end
      local.get 0
      local.get 4
      i32.store offset=4
      local.get 0
      local.get 1
      i32.store
      local.get 3
      i32.const 32
      i32.add
      global.set 0
      return
    end
    local.get 2
    i32.const 131072
    i32.const 14
    call 126
    unreachable)
  (func (;69;) (type 0) (param i32 i32)
    (local i32 i32 i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          call 142
          local.tee 1
          call 134
          br_if 0 (;@3;)
          i32.const 0
          local.set 3
          i32.const 0
          local.set 4
          i32.const 0
          local.set 5
          i32.const 0
          local.set 6
          br 1 (;@2;)
        end
        local.get 2
        i32.const 12
        i32.add
        local.get 1
        call 85
        local.get 2
        i32.const 12
        i32.add
        local.get 1
        call 123
        local.set 6
        local.get 2
        i32.const 12
        i32.add
        local.get 1
        call 123
        local.set 3
        local.get 2
        i32.const 12
        i32.add
        local.get 1
        call 123
        local.set 4
        local.get 2
        i32.const 12
        i32.add
        local.get 1
        call 123
        local.set 5
        local.get 2
        i32.load offset=16
        local.get 2
        i32.load offset=12
        i32.ne
        br_if 1 (;@1;)
        local.get 2
        i32.load8_u offset=28
        i32.eqz
        br_if 0 (;@2;)
        i32.const 0
        i32.const 0
        i32.store offset=142016
        i32.const 0
        i32.const 0
        i32.store8 offset=142020
      end
      local.get 0
      local.get 5
      i32.store offset=12
      local.get 0
      local.get 4
      i32.store offset=8
      local.get 0
      local.get 3
      i32.store offset=4
      local.get 0
      local.get 6
      i32.store
      local.get 2
      i32.const 32
      i32.add
      global.set 0
      return
    end
    local.get 1
    i32.const 131072
    i32.const 14
    call 126
    unreachable)
  (func (;70;) (type 15) (param i32 i32 i32 i32)
    local.get 0
    i32.const 131669
    i32.const 11
    local.get 1
    call 71
    local.set 0
    local.get 2
    call 143
    local.tee 1
    call 136
    local.get 3
    local.get 1
    call 136
    local.get 0
    local.get 1
    call 144)
  (func (;71;) (type 10) (param i32 i32 i32 i32) (result i32)
    local.get 0
    call 38
    local.tee 0
    local.get 1
    local.get 2
    call 12
    drop
    local.get 3
    local.get 0
    call 136
    local.get 0)
  (func (;72;) (type 5) (param i32)
    local.get 0
    i32.const 131960
    i32.const 0
    call 137)
  (func (;73;) (type 2) (param i32 i32) (result i32)
    local.get 0
    i32.const 131680
    i32.const 6
    local.get 1
    call 71
    call 145)
  (func (;74;) (type 0) (param i32 i32)
    (local i32 i32)
    local.get 0
    call 142
    local.set 2
    block  ;; label = @1
      local.get 1
      i32.load
      local.tee 3
      i32.eqz
      br_if 0 (;@1;)
      local.get 3
      call 143
      local.tee 0
      call 136
      local.get 1
      i32.load offset=4
      local.get 0
      call 136
      local.get 1
      i32.load offset=8
      local.get 0
      call 136
      local.get 1
      i32.load offset=12
      local.get 0
      call 136
      local.get 2
      local.get 0
      call 144
      return
    end
    local.get 2
    i32.const 131960
    i32.const 0
    call 137)
  (func (;75;) (type 2) (param i32 i32) (result i32)
    local.get 0
    call 38
    local.tee 0
    i32.const 131661
    i32.const 8
    call 12
    drop
    local.get 0
    local.get 1
    call 135
    local.get 0)
  (func (;76;) (type 14) (param i32 i32 i32)
    (local i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    i32.const 8
    i32.add
    local.get 1
    local.get 2
    call 77
    local.tee 2
    call 85
    local.get 3
    i32.const 8
    i32.add
    i32.const 20
    local.get 2
    call 124
    local.set 4
    local.get 3
    i32.const 8
    i32.add
    i32.const 32
    local.get 2
    call 124
    local.set 5
    local.get 3
    i32.const 8
    i32.add
    local.get 2
    call 88
    local.set 6
    local.get 3
    i32.const 8
    i32.add
    local.get 2
    call 87
    local.set 7
    local.get 3
    i32.const 8
    i32.add
    local.get 2
    call 139
    local.set 8
    local.get 3
    i32.const 0
    i32.store8 offset=31
    local.get 3
    i32.const 8
    i32.add
    local.get 3
    i32.const 31
    i32.add
    i32.const 1
    local.get 2
    call 140
    i32.const 2147483646
    local.set 1
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 3
          i32.load8_u offset=31
          br_table 2 (;@1;) 1 (;@2;) 0 (;@3;)
        end
        local.get 2
        i32.const 131734
        i32.const 13
        call 126
        unreachable
      end
      local.get 3
      i32.const 8
      i32.add
      local.get 2
      call 88
      local.set 1
    end
    block  ;; label = @1
      local.get 3
      i32.load offset=12
      local.get 3
      i32.load offset=8
      i32.ne
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 3
        i32.load8_u offset=24
        i32.eqz
        br_if 0 (;@2;)
        i32.const 0
        i32.const 0
        i32.store offset=142016
        i32.const 0
        i32.const 0
        i32.store8 offset=142020
      end
      local.get 0
      local.get 1
      i32.store offset=24
      local.get 0
      local.get 7
      i32.store offset=20
      local.get 0
      local.get 6
      i32.store offset=16
      local.get 0
      local.get 5
      i32.store offset=12
      local.get 0
      local.get 4
      i32.store offset=8
      local.get 0
      local.get 8
      i64.store
      local.get 3
      i32.const 32
      i32.add
      global.set 0
      return
    end
    local.get 2
    i32.const 131072
    i32.const 14
    call 126
    unreachable)
  (func (;77;) (type 2) (param i32 i32) (result i32)
    local.get 0
    call 38
    local.tee 0
    i32.const 131654
    i32.const 7
    call 12
    drop
    local.get 0
    local.get 1
    call 135
    local.get 0)
  (func (;78;) (type 5) (param i32)
    i32.const -20
    i32.const 131960
    i32.const 0
    call 20
    drop
    local.get 0
    i32.const -20
    call 27
    drop)
  (func (;79;) (type 4) (param i32) (result i32)
    (local i32)
    i32.const 131796
    i32.const 16
    call 97
    local.tee 1
    local.get 0
    call 135
    local.get 1)
  (func (;80;) (type 15) (param i32 i32 i32 i32)
    (local i32 i64)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 4
    global.set 0
    i64.const 0
    local.set 5
    block  ;; label = @1
      local.get 1
      local.get 3
      call 141
      i32.eqz
      br_if 0 (;@1;)
      local.get 4
      local.get 2
      local.get 3
      call 76
      local.get 0
      i32.const 32
      i32.add
      local.get 4
      i32.const 24
      i32.add
      i64.load
      i64.store
      local.get 0
      i32.const 24
      i32.add
      local.get 4
      i32.const 16
      i32.add
      i64.load
      i64.store
      local.get 0
      i32.const 16
      i32.add
      local.get 4
      i32.const 8
      i32.add
      i64.load
      i64.store
      local.get 0
      local.get 4
      i64.load
      i64.store offset=8
      i64.const 1
      local.set 5
    end
    local.get 0
    local.get 5
    i64.store
    local.get 4
    i32.const 32
    i32.add
    global.set 0)
  (func (;81;) (type 0) (param i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        br_if 0 (;@2;)
        local.get 1
        call 82
        i32.eqz
        br_if 1 (;@1;)
        call 83
        local.get 1
        call 84
        return
      end
      call 83
      call 78
      return
    end
    i32.const 131377
    i32.const 30
    call 59
    unreachable)
  (func (;82;) (type 4) (param i32) (result i32)
    local.get 0
    i32.const 142029
    call 8
    drop
    i32.const 142029
    call 9
    i32.const 0
    i32.gt_s)
  (func (;83;) (type 1) (result i32)
    i32.const 131832
    i32.const 20
    call 97)
  (func (;84;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call 138)
  (func (;85;) (type 0) (param i32 i32)
    (local i32)
    local.get 1
    call 86
    local.tee 2
    call 6
    local.set 1
    local.get 0
    i32.const 0
    i32.store8 offset=16
    local.get 0
    local.get 1
    i32.store offset=12
    local.get 0
    local.get 2
    i32.store offset=8
    local.get 0
    local.get 1
    i32.store offset=4
    local.get 0
    i32.const 0
    i32.store)
  (func (;86;) (type 4) (param i32) (result i32)
    (local i32)
    local.get 0
    call 55
    local.tee 1
    call 26
    drop
    local.get 1)
  (func (;87;) (type 2) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 88
    call 89)
  (func (;88;) (type 2) (param i32 i32) (result i32)
    local.get 0
    local.get 0
    local.get 1
    call 123
    local.get 1
    call 124)
  (func (;89;) (type 4) (param i32) (result i32)
    (local i32)
    local.get 0
    call 55
    local.tee 1
    call 24
    drop
    local.get 1)
  (func (;90;) (type 0) (param i32 i32)
    local.get 0
    call 91
    local.get 1
    call 92)
  (func (;91;) (type 4) (param i32) (result i32)
    (local i32)
    call 55
    local.tee 1
    local.get 0
    call 23
    drop
    local.get 1)
  (func (;92;) (type 0) (param i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    call 120
    local.tee 3
    i32.const 24
    i32.shl
    local.get 3
    i32.const 65280
    i32.and
    i32.const 8
    i32.shl
    i32.or
    local.get 3
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 3
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    i32.store offset=12
    local.get 1
    local.get 2
    i32.const 12
    i32.add
    i32.const 4
    call 127
    local.get 1
    local.get 0
    call 128
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;93;) (type 1) (result i32)
    (local i32)
    call 55
    local.tee 0
    call 7
    local.get 0)
  (func (;94;) (type 3)
    (local i32)
    call 55
    local.tee 0
    call 10
    block  ;; label = @1
      local.get 0
      call 93
      call 39
      i32.eqz
      br_if 0 (;@1;)
      return
    end
    i32.const 131960
    i32.const 36
    call 0
    unreachable)
  (func (;95;) (type 4) (param i32) (result i32)
    (local i32)
    local.get 0
    call 55
    local.tee 1
    call 11
    drop
    local.get 1)
  (func (;96;) (type 15) (param i32 i32 i32 i32)
    (local i32)
    i32.const 131452
    i32.const 23
    call 97
    local.tee 4
    local.get 0
    local.get 1
    call 12
    drop
    local.get 4
    i32.const 131475
    i32.const 3
    call 12
    drop
    local.get 4
    local.get 2
    local.get 3
    call 12
    drop
    local.get 4
    call 13
    unreachable)
  (func (;97;) (type 2) (param i32 i32) (result i32)
    (local i32)
    call 55
    local.tee 2
    local.get 0
    local.get 1
    call 20
    drop
    local.get 2)
  (func (;98;) (type 0) (param i32 i32)
    (local i64 i32 i32 i32)
    block  ;; label = @1
      local.get 1
      i32.const 131140
      i32.const 6
      call 99
      call 14
      local.tee 2
      i64.const 4294967296
      i64.ge_u
      br_if 0 (;@1;)
      i32.const 0
      local.set 3
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          i32.wrap_i64
          local.tee 4
          br_if 0 (;@3;)
          br 1 (;@2;)
        end
        local.get 4
        i32.const 28523
        i32.eq
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 1
          i32.load
          i32.const 0
          i32.load offset=142024
          i32.lt_s
          br_if 0 (;@3;)
          i32.const 1
          local.set 3
          call 62
          local.set 5
          br 1 (;@2;)
        end
        i32.const 1
        local.set 3
        local.get 1
        i32.const 131140
        i32.const 6
        call 99
        call 95
        local.set 5
      end
      local.get 0
      local.get 5
      i32.store offset=8
      local.get 0
      local.get 4
      i32.store offset=4
      local.get 0
      local.get 3
      i32.store
      return
    end
    i32.const 131140
    i32.const 6
    i32.const 131072
    i32.const 14
    call 96
    unreachable)
  (func (;99;) (type 6) (param i32 i32 i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 3
      i32.const 0
      i32.load offset=142024
      i32.lt_s
      br_if 0 (;@1;)
      local.get 1
      local.get 2
      i32.const 131478
      i32.const 17
      call 96
      unreachable
    end
    local.get 0
    local.get 3
    i32.const 1
    i32.add
    i32.store
    local.get 3)
  (func (;100;) (type 4) (param i32) (result i32)
    (local i32 i32 i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 0
    i32.load offset=8
    local.set 2
    local.get 1
    i32.const 0
    i32.store offset=4
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          i32.load
          local.get 2
          i32.const 2
          i32.shl
          local.get 1
          i32.const 4
          i32.add
          i32.const 4
          call 101
          br_if 0 (;@3;)
          local.get 1
          i32.load offset=4
          local.set 3
          local.get 0
          local.get 2
          i32.const 1
          i32.add
          i32.store offset=8
          local.get 3
          i32.const 24
          i32.shl
          local.get 3
          i32.const 65280
          i32.and
          i32.const 8
          i32.shl
          i32.or
          local.get 3
          i32.const 8
          i32.shr_u
          i32.const 65280
          i32.and
          local.get 3
          i32.const 24
          i32.shr_u
          i32.or
          i32.or
          call 38
          local.set 3
          local.get 1
          i64.const 0
          i64.store offset=8
          local.get 3
          call 6
          local.tee 0
          i32.const 9
          i32.ge_u
          br_if 1 (;@2;)
          local.get 3
          i32.const 0
          local.get 1
          i32.const 8
          i32.add
          local.get 0
          i32.sub
          i32.const 8
          i32.add
          local.get 0
          call 101
          drop
          local.get 1
          i64.load offset=8
          local.tee 4
          i64.const 56
          i64.shl
          local.get 4
          i64.const 65280
          i64.and
          i64.const 40
          i64.shl
          i64.or
          local.get 4
          i64.const 16711680
          i64.and
          i64.const 24
          i64.shl
          local.get 4
          i64.const 4278190080
          i64.and
          i64.const 8
          i64.shl
          i64.or
          i64.or
          local.get 4
          i64.const 8
          i64.shr_u
          i64.const 4278190080
          i64.and
          local.get 4
          i64.const 24
          i64.shr_u
          i64.const 16711680
          i64.and
          i64.or
          local.get 4
          i64.const 40
          i64.shr_u
          i64.const 65280
          i64.and
          local.get 4
          i64.const 56
          i64.shr_u
          i64.or
          i64.or
          i64.or
          local.tee 4
          i64.const 4294967296
          i64.ge_u
          br_if 2 (;@1;)
          local.get 1
          i32.const 16
          i32.add
          global.set 0
          local.get 4
          i32.wrap_i64
          return
        end
        i32.const 131135
        i32.const 5
        i32.const 131478
        i32.const 17
        call 96
        unreachable
      end
      i32.const 131135
      i32.const 5
      i32.const 131072
      i32.const 14
      call 96
      unreachable
    end
    i32.const 131135
    i32.const 5
    i32.const 131072
    i32.const 14
    call 96
    unreachable)
  (func (;101;) (type 10) (param i32 i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 3
    local.get 2
    call 19
    i32.const 0
    i32.ne)
  (func (;102;) (type 15) (param i32 i32 i32 i32)
    (local i32)
    i32.const 0
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.load
        i32.const 0
        i32.load offset=142024
        i32.lt_s
        br_if 0 (;@2;)
        i32.const 1
        local.set 4
        br 1 (;@1;)
      end
      local.get 1
      local.get 2
      local.get 3
      call 99
      call 95
      local.tee 1
      call 6
      i32.const 32
      i32.eq
      br_if 0 (;@1;)
      local.get 2
      local.get 3
      i32.const 131718
      i32.const 16
      call 96
      unreachable
    end
    local.get 0
    local.get 1
    i32.store offset=4
    local.get 0
    local.get 4
    i32.store)
  (func (;103;) (type 5) (param i32)
    (local i32 i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 1
    global.set 0
    i32.const 0
    call 95
    local.tee 2
    call 6
    local.set 3
    local.get 1
    i32.const 0
    i32.store8 offset=20
    local.get 1
    local.get 3
    i32.store offset=16
    local.get 1
    local.get 2
    i32.store offset=12
    local.get 1
    local.get 3
    i32.store offset=8
    local.get 1
    i32.const 0
    i32.store offset=4
    local.get 1
    i32.const 4
    i32.add
    i32.const 20
    call 104
    local.set 2
    local.get 1
    i32.const 4
    i32.add
    i32.const 32
    call 104
    local.set 4
    local.get 1
    i32.const 4
    i32.add
    call 105
    local.set 5
    local.get 1
    i32.const 4
    i32.add
    call 105
    call 89
    local.set 6
    local.get 1
    i64.const 0
    i64.store offset=24
    local.get 1
    i32.const 4
    i32.add
    local.get 1
    i32.const 24
    i32.add
    i32.const 8
    call 106
    local.get 1
    i64.load offset=24
    local.set 7
    local.get 1
    i32.const 0
    i32.store8 offset=24
    local.get 1
    i32.const 4
    i32.add
    local.get 1
    i32.const 24
    i32.add
    i32.const 1
    call 106
    i32.const 2147483646
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.load8_u offset=24
          br_table 2 (;@1;) 1 (;@2;) 0 (;@3;)
        end
        i32.const 131129
        i32.const 6
        i32.const 131734
        i32.const 13
        call 96
        unreachable
      end
      local.get 1
      i32.const 4
      i32.add
      call 105
      local.set 3
    end
    block  ;; label = @1
      local.get 1
      i32.load offset=8
      local.get 1
      i32.load offset=4
      i32.ne
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 1
        i32.load8_u offset=20
        i32.eqz
        br_if 0 (;@2;)
        i32.const 0
        i32.const 0
        i32.store offset=142016
        i32.const 0
        i32.const 0
        i32.store8 offset=142020
      end
      local.get 0
      local.get 3
      i32.store offset=24
      local.get 0
      local.get 6
      i32.store offset=20
      local.get 0
      local.get 5
      i32.store offset=16
      local.get 0
      local.get 4
      i32.store offset=12
      local.get 0
      local.get 2
      i32.store offset=8
      local.get 0
      local.get 7
      i64.const 56
      i64.shl
      local.get 7
      i64.const 65280
      i64.and
      i64.const 40
      i64.shl
      i64.or
      local.get 7
      i64.const 16711680
      i64.and
      i64.const 24
      i64.shl
      local.get 7
      i64.const 4278190080
      i64.and
      i64.const 8
      i64.shl
      i64.or
      i64.or
      local.get 7
      i64.const 8
      i64.shr_u
      i64.const 4278190080
      i64.and
      local.get 7
      i64.const 24
      i64.shr_u
      i64.const 16711680
      i64.and
      i64.or
      local.get 7
      i64.const 40
      i64.shr_u
      i64.const 65280
      i64.and
      local.get 7
      i64.const 56
      i64.shr_u
      i64.or
      i64.or
      i64.or
      i64.store
      local.get 1
      i32.const 32
      i32.add
      global.set 0
      return
    end
    i32.const 131129
    i32.const 6
    i32.const 131072
    i32.const 14
    call 96
    unreachable)
  (func (;104;) (type 2) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 8
    i32.add
    local.get 0
    i32.load offset=8
    local.get 0
    i32.load
    local.tee 3
    local.get 1
    call 125
    block  ;; label = @1
      local.get 2
      i32.load offset=8
      i32.const 1
      i32.ne
      br_if 0 (;@1;)
      local.get 2
      i32.load offset=12
      local.set 4
      local.get 0
      local.get 3
      local.get 1
      i32.add
      i32.store
      local.get 2
      i32.const 16
      i32.add
      global.set 0
      local.get 4
      return
    end
    i32.const 131129
    i32.const 6
    i32.const 131610
    i32.const 15
    call 96
    unreachable)
  (func (;105;) (type 4) (param i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 1
    i32.const 12
    i32.add
    i32.const 4
    call 106
    local.get 0
    local.get 1
    i32.load offset=12
    local.tee 2
    i32.const 24
    i32.shl
    local.get 2
    i32.const 65280
    i32.and
    i32.const 8
    i32.shl
    i32.or
    local.get 2
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 2
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    call 104
    local.set 2
    local.get 1
    i32.const 16
    i32.add
    global.set 0
    local.get 2)
  (func (;106;) (type 14) (param i32 i32 i32)
    block  ;; label = @1
      local.get 0
      i32.const 8
      i32.add
      local.get 0
      i32.load
      local.get 1
      local.get 2
      call 132
      i32.eqz
      br_if 0 (;@1;)
      i32.const 131129
      i32.const 6
      call 154
      unreachable
    end
    local.get 0
    local.get 0
    i32.load
    local.get 2
    i32.add
    i32.store)
  (func (;107;) (type 1) (result i32)
    (local i64)
    block  ;; label = @1
      i32.const 0
      call 14
      local.tee 0
      i64.const 4294967296
      i64.lt_u
      br_if 0 (;@1;)
      i32.const 131135
      i32.const 5
      i32.const 131072
      i32.const 14
      call 96
      unreachable
    end
    local.get 0
    i32.wrap_i64)
  (func (;108;) (type 12) (result i64)
    i32.const 1
    call 14)
  (func (;109;) (type 5) (param i32)
    block  ;; label = @1
      i32.const 0
      i32.load offset=142024
      local.get 0
      i32.gt_s
      br_if 0 (;@1;)
      return
    end
    i32.const 131495
    i32.const 18
    call 0
    unreachable)
  (func (;110;) (type 0) (param i32 i32)
    block  ;; label = @1
      local.get 1
      local.get 0
      i32.lt_u
      br_if 0 (;@1;)
      return
    end
    i32.const 131495
    i32.const 18
    call 0
    unreachable)
  (func (;111;) (type 5) (param i32)
    block  ;; label = @1
      call 15
      local.get 0
      i32.ne
      br_if 0 (;@1;)
      return
    end
    i32.const 131513
    i32.const 25
    call 0
    unreachable)
  (func (;112;) (type 3)
    block  ;; label = @1
      i32.const 0
      i32.load offset=142024
      i32.const 0
      i32.lt_s
      br_if 0 (;@1;)
      return
    end
    i32.const 131478
    i32.const 17
    call 0
    unreachable)
  (func (;113;) (type 3)
    i32.const 0
    call 15
    i32.store offset=142024)
  (func (;114;) (type 5) (param i32)
    local.get 0
    call 43
    call 16
    drop)
  (func (;115;) (type 20) (param i32 i32 i32 i64)
    (local i32)
    call 62
    local.tee 4
    local.get 3
    call 121
    local.get 2
    local.get 4
    call 63
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;116;) (type 15) (param i32 i32 i32 i32)
    call 62
    drop
    local.get 2
    local.get 3
    i32.load
    call 91
    call 63
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;117;) (type 16) (param i32 i32 i32 i32 i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 5
    global.set 0
    block  ;; label = @1
      local.get 3
      call 118
      br_if 0 (;@1;)
      local.get 2
      local.get 3
      call 119
      local.get 4
      call 120
      local.set 6
      i32.const 0
      local.set 3
      loop  ;; label = @2
        local.get 3
        i32.const 4
        i32.add
        local.tee 7
        local.get 6
        i32.gt_u
        br_if 1 (;@1;)
        local.get 5
        i32.const 0
        i32.store offset=12
        local.get 4
        local.get 3
        local.get 5
        i32.const 12
        i32.add
        i32.const 4
        call 101
        drop
        local.get 2
        local.get 5
        i32.load offset=12
        local.tee 3
        i32.const 24
        i32.shl
        local.get 3
        i32.const 65280
        i32.and
        i32.const 8
        i32.shl
        i32.or
        local.get 3
        i32.const 8
        i32.shr_u
        i32.const 65280
        i32.and
        local.get 3
        i32.const 24
        i32.shr_u
        i32.or
        i32.or
        call 119
        local.get 7
        local.set 3
        br 0 (;@2;)
      end
    end
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store
    local.get 5
    i32.const 16
    i32.add
    global.set 0)
  (func (;118;) (type 4) (param i32) (result i32)
    local.get 0
    call 120
    i32.eqz)
  (func (;119;) (type 0) (param i32 i32)
    call 62
    drop
    local.get 0
    local.get 1
    call 38
    call 63)
  (func (;120;) (type 4) (param i32) (result i32)
    local.get 0
    call 6)
  (func (;121;) (type 11) (param i32 i64)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i64.const 0
    i64.store offset=8
    local.get 2
    local.get 1
    i32.const 0
    local.get 2
    i32.const 8
    i32.add
    call 159
    local.get 0
    local.get 2
    i32.load
    local.get 2
    i32.load offset=4
    call 20
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;122;) (type 16) (param i32 i32 i32 i32 i32)
    (local i32 i32)
    local.get 1
    local.get 2
    local.get 3
    local.get 4
    local.get 0
    i32.load
    local.tee 5
    local.get 0
    i32.load offset=4
    local.tee 6
    local.get 5
    local.get 6
    local.get 0
    i32.load offset=8
    i64.load
    local.get 0
    i32.load offset=12
    i64.load
    local.get 0
    i32.load offset=16
    i32.load
    call 18
    drop)
  (func (;123;) (type 2) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 2
    i32.const 12
    i32.add
    i32.const 4
    local.get 1
    call 140
    local.get 2
    i32.load offset=12
    local.set 1
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 1
    i32.const 24
    i32.shl
    local.get 1
    i32.const 65280
    i32.and
    i32.const 8
    i32.shl
    i32.or
    local.get 1
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 1
    i32.const 24
    i32.shr_u
    i32.or
    i32.or)
  (func (;124;) (type 6) (param i32 i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    i32.const 8
    i32.add
    local.get 0
    i32.load offset=8
    local.get 0
    i32.load
    local.tee 4
    local.get 1
    call 125
    block  ;; label = @1
      local.get 3
      i32.load offset=8
      i32.const 1
      i32.ne
      br_if 0 (;@1;)
      local.get 3
      i32.load offset=12
      local.set 2
      local.get 0
      local.get 4
      local.get 1
      i32.add
      i32.store
      local.get 3
      i32.const 16
      i32.add
      global.set 0
      local.get 2
      return
    end
    local.get 2
    i32.const 131610
    i32.const 15
    call 126
    unreachable)
  (func (;125;) (type 15) (param i32 i32 i32 i32)
    (local i32)
    local.get 1
    local.get 2
    local.get 3
    call 1
    local.tee 4
    call 22
    local.set 3
    local.get 0
    local.get 4
    i32.store offset=4
    local.get 0
    local.get 3
    i32.eqz
    i32.store)
  (func (;126;) (type 14) (param i32 i32 i32)
    (local i32)
    i32.const 131691
    i32.const 27
    call 97
    local.tee 3
    local.get 0
    call 2
    drop
    local.get 3
    i32.const 131475
    i32.const 3
    call 12
    drop
    local.get 3
    local.get 1
    local.get 2
    call 12
    drop
    local.get 3
    call 13
    unreachable)
  (func (;127;) (type 14) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 12
    drop)
  (func (;128;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call 2
    drop)
  (func (;129;) (type 2) (param i32 i32) (result i32)
    (local i32 i32)
    local.get 1
    i32.load
    local.tee 2
    i32.const 4
    i32.add
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const -5
        i32.gt_u
        br_if 0 (;@2;)
        local.get 3
        i32.const 16
        i32.gt_u
        br_if 1 (;@1;)
        local.get 1
        local.get 3
        i32.store
        local.get 0
        local.get 2
        i32.add
        i32.load align=1
        local.tee 2
        i32.const 24
        i32.shl
        local.get 2
        i32.const 65280
        i32.and
        i32.const 8
        i32.shl
        i32.or
        local.get 2
        i32.const 8
        i32.shr_u
        i32.const 65280
        i32.and
        local.get 2
        i32.const 24
        i32.shr_u
        i32.or
        i32.or
        return
      end
      local.get 2
      local.get 3
      call 131
      unreachable
    end
    local.get 3
    i32.const 16
    call 131
    unreachable)
  (func (;130;) (type 18) (param i32 i32) (result i64)
    (local i32 i32 i64)
    local.get 1
    i32.load
    local.tee 2
    i32.const 8
    i32.add
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const -9
        i32.gt_u
        br_if 0 (;@2;)
        local.get 3
        i32.const 16
        i32.gt_u
        br_if 1 (;@1;)
        local.get 1
        local.get 3
        i32.store
        local.get 0
        local.get 2
        i32.add
        i64.load align=1
        local.tee 4
        i64.const 56
        i64.shl
        local.get 4
        i64.const 65280
        i64.and
        i64.const 40
        i64.shl
        i64.or
        local.get 4
        i64.const 16711680
        i64.and
        i64.const 24
        i64.shl
        local.get 4
        i64.const 4278190080
        i64.and
        i64.const 8
        i64.shl
        i64.or
        i64.or
        local.get 4
        i64.const 8
        i64.shr_u
        i64.const 4278190080
        i64.and
        local.get 4
        i64.const 24
        i64.shr_u
        i64.const 16711680
        i64.and
        i64.or
        local.get 4
        i64.const 40
        i64.shr_u
        i64.const 65280
        i64.and
        local.get 4
        i64.const 56
        i64.shr_u
        i64.or
        i64.or
        i64.or
        return
      end
      local.get 2
      local.get 3
      call 131
      unreachable
    end
    local.get 3
    i32.const 16
    call 131
    unreachable)
  (func (;131;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call 189
    unreachable)
  (func (;132;) (type 10) (param i32 i32 i32 i32) (result i32)
    (local i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              i32.load8_u offset=8
              br_if 0 (;@5;)
              local.get 0
              i32.load
              local.tee 4
              call 120
              local.tee 5
              i32.const 10000
              i32.gt_u
              br_if 1 (;@4;)
              i32.const 0
              i32.load8_u offset=142020
              i32.const 255
              i32.and
              br_if 1 (;@4;)
              i32.const 0
              local.get 5
              i32.store offset=142016
              i32.const 0
              i32.const 1
              i32.store8 offset=142020
              local.get 4
              i32.const 0
              i32.const 132016
              local.get 5
              call 101
              drop
              local.get 0
              i32.const 1
              i32.store8 offset=8
            end
            i32.const 1
            local.set 4
            local.get 3
            local.get 1
            i32.add
            local.tee 0
            i32.const 0
            i32.load offset=142016
            i32.gt_u
            br_if 1 (;@3;)
            local.get 0
            local.get 1
            i32.lt_u
            br_if 2 (;@2;)
            local.get 0
            i32.const 10000
            i32.gt_u
            br_if 3 (;@1;)
            local.get 2
            local.get 3
            local.get 1
            i32.const 132016
            i32.add
            local.get 3
            call 133
            i32.const 0
            return
          end
          local.get 0
          i32.const 0
          i32.store8 offset=8
          local.get 4
          local.get 1
          local.get 2
          local.get 3
          call 101
          local.set 4
        end
        local.get 4
        return
      end
      local.get 1
      local.get 0
      call 131
      unreachable
    end
    local.get 0
    i32.const 10000
    call 131
    unreachable)
  (func (;133;) (type 15) (param i32 i32 i32 i32)
    block  ;; label = @1
      local.get 1
      local.get 3
      i32.eq
      br_if 0 (;@1;)
      local.get 1
      local.get 3
      call 190
      unreachable
    end
    local.get 0
    local.get 2
    local.get 1
    call 192
    drop)
  (func (;134;) (type 4) (param i32) (result i32)
    local.get 0
    i32.const -25
    call 26
    drop
    i32.const -25
    call 6)
  (func (;135;) (type 0) (param i32 i32)
    local.get 1
    local.get 0
    call 136)
  (func (;136;) (type 0) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    i32.const 24
    i32.shl
    local.get 0
    i32.const 65280
    i32.and
    i32.const 8
    i32.shl
    i32.or
    local.get 0
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 0
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    i32.store offset=12
    local.get 1
    local.get 2
    i32.const 12
    i32.add
    i32.const 4
    call 127
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;137;) (type 14) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 97
    call 27
    drop)
  (func (;138;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call 27
    drop)
  (func (;139;) (type 18) (param i32 i32) (result i64)
    (local i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i64.const 0
    i64.store offset=8
    local.get 0
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    local.get 1
    call 140
    local.get 2
    i64.load offset=8
    local.set 3
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 3
    i64.const 56
    i64.shl
    local.get 3
    i64.const 65280
    i64.and
    i64.const 40
    i64.shl
    i64.or
    local.get 3
    i64.const 16711680
    i64.and
    i64.const 24
    i64.shl
    local.get 3
    i64.const 4278190080
    i64.and
    i64.const 8
    i64.shl
    i64.or
    i64.or
    local.get 3
    i64.const 8
    i64.shr_u
    i64.const 4278190080
    i64.and
    local.get 3
    i64.const 24
    i64.shr_u
    i64.const 16711680
    i64.and
    i64.or
    local.get 3
    i64.const 40
    i64.shr_u
    i64.const 65280
    i64.and
    local.get 3
    i64.const 56
    i64.shr_u
    i64.or
    i64.or
    i64.or)
  (func (;140;) (type 15) (param i32 i32 i32 i32)
    block  ;; label = @1
      local.get 0
      i32.const 8
      i32.add
      local.get 0
      i32.load
      local.get 1
      local.get 2
      call 132
      i32.eqz
      br_if 0 (;@1;)
      local.get 3
      call 155
      unreachable
    end
    local.get 0
    local.get 0
    i32.load
    local.get 2
    i32.add
    i32.store)
  (func (;141;) (type 2) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 67
    i32.const 0
    i32.ne)
  (func (;142;) (type 4) (param i32) (result i32)
    local.get 0
    call 38
    local.tee 0
    i32.const 131686
    i32.const 5
    call 12
    drop
    local.get 0)
  (func (;143;) (type 1) (result i32)
    i32.const 131960
    i32.const 0
    call 97)
  (func (;144;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call 27
    drop)
  (func (;145;) (type 4) (param i32) (result i32)
    (local i64)
    block  ;; label = @1
      local.get 0
      local.get 0
      call 61
      local.tee 1
      i64.const 4294967296
      i64.lt_u
      br_if 0 (;@1;)
      local.get 0
      i32.const 131072
      i32.const 14
      call 126
      unreachable
    end
    local.get 1
    i32.wrap_i64)
  (func (;146;) (type 4) (param i32) (result i32)
    (local i64)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          local.get 0
          call 61
          local.tee 1
          i64.const 1
          i64.gt_u
          br_if 0 (;@3;)
          i32.const 0
          local.set 0
          local.get 1
          i32.wrap_i64
          br_table 2 (;@1;) 1 (;@2;) 2 (;@1;)
        end
        local.get 0
        i32.const 131747
        i32.const 18
        call 126
        unreachable
      end
      i32.const 1
      local.set 0
    end
    local.get 0)
  (func (;147;) (type 2) (param i32 i32) (result i32)
    (local i32)
    call 62
    local.tee 2
    local.get 0
    local.get 1
    call 97
    call 63
    local.get 2)
  (func (;148;) (type 0) (param i32 i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    call 149
    local.get 2
    i32.load offset=12
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.load offset=8
        local.tee 4
        i32.eqz
        br_if 0 (;@2;)
        local.get 3
        local.set 5
        br 1 (;@1;)
      end
      local.get 2
      local.get 1
      i32.load offset=8
      local.get 1
      i32.load
      local.tee 4
      local.get 3
      call 125
      block  ;; label = @2
        local.get 2
        i32.load
        i32.const 1
        i32.eq
        br_if 0 (;@2;)
        i32.const 15
        local.set 5
        i32.const 131610
        local.set 4
        br 1 (;@1;)
      end
      local.get 2
      i32.load offset=4
      local.set 5
      local.get 1
      local.get 4
      local.get 3
      i32.add
      i32.store
      i32.const 0
      local.set 4
    end
    local.get 0
    local.get 5
    i64.extend_i32_u
    i64.const 32
    i64.shl
    local.get 4
    i64.extend_i32_u
    i64.or
    i64.store align=4
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;149;) (type 0) (param i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i32.const 0
    i32.store offset=12
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        local.get 2
        i32.const 12
        i32.add
        i32.const 4
        call 157
        local.tee 1
        br_if 0 (;@2;)
        local.get 2
        i32.load offset=12
        local.tee 3
        i32.const 24
        i32.shl
        local.get 3
        i32.const 65280
        i32.and
        i32.const 8
        i32.shl
        i32.or
        local.get 3
        i32.const 8
        i32.shr_u
        i32.const 65280
        i32.and
        local.get 3
        i32.const 24
        i32.shr_u
        i32.or
        i32.or
        local.set 3
        br 1 (;@1;)
      end
      i32.const 15
      local.set 3
    end
    local.get 0
    local.get 1
    i32.store
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;150;) (type 0) (param i32 i32)
    block  ;; label = @1
      local.get 0
      i32.const 2147483646
      i32.ne
      br_if 0 (;@1;)
      local.get 1
      i32.const 0
      call 151
      return
    end
    local.get 1
    i32.const 1
    call 151
    local.get 0
    local.get 1
    call 92)
  (func (;151;) (type 0) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.store8 offset=15
    local.get 0
    local.get 2
    i32.const 15
    i32.add
    i32.const 1
    call 127
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;152;) (type 0) (param i32 i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.load offset=4
        local.tee 3
        i32.const 4
        i32.add
        local.tee 4
        local.get 1
        i32.load offset=8
        i32.le_u
        br_if 0 (;@2;)
        i32.const 0
        local.set 1
        br 1 (;@1;)
      end
      local.get 1
      i32.load
      local.set 5
      local.get 2
      i32.const 0
      i32.store offset=12
      local.get 5
      i32.load
      local.get 3
      local.get 2
      i32.const 12
      i32.add
      i32.const 4
      call 101
      drop
      local.get 2
      i32.load offset=12
      local.set 3
      local.get 1
      local.get 4
      i32.store offset=4
      local.get 3
      i32.const 24
      i32.shl
      local.get 3
      i32.const 65280
      i32.and
      i32.const 8
      i32.shl
      i32.or
      local.get 3
      i32.const 8
      i32.shr_u
      i32.const 65280
      i32.and
      local.get 3
      i32.const 24
      i32.shr_u
      i32.or
      i32.or
      local.set 3
      i32.const 1
      local.set 1
    end
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;153;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call 97
    local.tee 1
    i32.const 131610
    i32.const 15
    call 12
    drop
    local.get 1
    call 13
    unreachable)
  (func (;154;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    i32.const 131610
    i32.const 15
    call 96
    unreachable)
  (func (;155;) (type 5) (param i32)
    local.get 0
    i32.const 131610
    i32.const 15
    call 126
    unreachable)
  (func (;156;) (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call 153
    unreachable)
  (func (;157;) (type 6) (param i32 i32 i32) (result i32)
    (local i32)
    i32.const 131610
    local.set 3
    block  ;; label = @1
      local.get 0
      i32.const 8
      i32.add
      local.get 0
      i32.load
      local.get 1
      local.get 2
      call 132
      br_if 0 (;@1;)
      local.get 0
      local.get 0
      i32.load
      local.get 2
      i32.add
      i32.store
      i32.const 0
      local.set 3
    end
    local.get 3)
  (func (;158;) (type 11) (param i32 i64)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    i64.const 0
    i64.store offset=8
    local.get 2
    local.get 1
    i32.const 0
    local.get 2
    i32.const 8
    i32.add
    call 159
    local.get 0
    local.get 2
    i32.load
    local.get 2
    i32.load offset=4
    call 137
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;159;) (type 21) (param i32 i64 i32 i32)
    (local i64 i64 i32 i32 i32 i32)
    local.get 3
    local.get 1
    i64.const 56
    i64.shl
    local.get 1
    i64.const 65280
    i64.and
    i64.const 40
    i64.shl
    i64.or
    local.get 1
    i64.const 16711680
    i64.and
    i64.const 24
    i64.shl
    local.get 1
    i64.const 4278190080
    i64.and
    i64.const 8
    i64.shl
    i64.or
    i64.or
    local.get 1
    i64.const 8
    i64.shr_u
    i64.const 4278190080
    i64.and
    local.get 1
    i64.const 24
    i64.shr_u
    i64.const 16711680
    i64.and
    i64.or
    local.get 1
    i64.const 40
    i64.shr_u
    local.tee 4
    i64.const 65280
    i64.and
    local.get 1
    i64.const 56
    i64.shr_u
    local.tee 5
    i64.or
    i64.or
    i64.or
    i64.store align=1
    local.get 0
    i32.const 8
    i32.const 0
    local.get 1
    i64.const 0
    i64.lt_s
    local.tee 6
    local.get 2
    i32.and
    i32.sub
    i32.const 255
    i32.and
    local.tee 7
    local.get 5
    i32.wrap_i64
    i32.eq
    local.tee 8
    i32.const 0
    local.get 7
    local.get 1
    i64.const 48
    i64.shr_u
    i32.wrap_i64
    i32.const 255
    i32.and
    i32.eq
    select
    local.tee 9
    local.get 8
    i32.add
    local.get 9
    i32.const 0
    local.get 7
    local.get 4
    i32.wrap_i64
    i32.const 255
    i32.and
    i32.eq
    select
    local.tee 8
    i32.add
    local.get 8
    i32.const 0
    local.get 7
    local.get 1
    i64.const 32
    i64.shr_u
    i32.wrap_i64
    i32.const 255
    i32.and
    i32.eq
    select
    local.tee 9
    i32.add
    local.get 9
    i32.const 0
    local.get 7
    local.get 1
    i32.wrap_i64
    local.tee 8
    i32.const 24
    i32.shr_u
    i32.eq
    select
    local.tee 9
    i32.add
    local.get 9
    i32.const 0
    local.get 7
    local.get 8
    i32.const 16
    i32.shr_u
    i32.const 255
    i32.and
    i32.eq
    select
    local.tee 9
    i32.add
    local.get 9
    i32.const 0
    local.get 7
    local.get 8
    i32.const 8
    i32.shr_u
    i32.const 255
    i32.and
    i32.eq
    select
    local.tee 7
    i32.add
    local.get 7
    i32.const 0
    local.get 1
    i64.eqz
    select
    i32.add
    local.tee 7
    local.get 7
    i32.const 0
    i32.ne
    local.get 6
    local.get 3
    local.get 7
    i32.const 7
    i32.and
    i32.add
    i32.load8_s
    i32.const 0
    i32.lt_s
    i32.xor
    i32.and
    local.get 2
    i32.and
    i32.sub
    local.tee 7
    i32.sub
    i32.store offset=4
    local.get 0
    local.get 3
    local.get 7
    i32.add
    i32.store)
  (func (;160;) (type 5) (param i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    call 161
    local.set 2
    local.get 1
    i64.const 0
    i64.store offset=8
    local.get 1
    local.get 0
    i64.extend_i32_u
    i32.const 1
    local.get 1
    i32.const 8
    i32.add
    call 159
    local.get 2
    local.get 1
    i32.load
    local.get 1
    i32.load offset=4
    call 137
    local.get 1
    i32.const 16
    i32.add
    global.set 0)
  (func (;161;) (type 1) (result i32)
    i32.const 131923
    i32.const 19
    call 97)
  (func (;162;) (type 3)
    block  ;; label = @1
      call 161
      call 146
      i32.eqz
      br_if 0 (;@1;)
      i32.const 131765
      i32.const 18
      call 59
      unreachable
    end)
  (func (;163;) (type 0) (param i32 i32)
    (local i32)
    call 62
    local.tee 2
    local.get 1
    i64.extend_i32_u
    call 121
    local.get 0
    local.get 2
    call 63)
  (func (;164;) (type 1) (result i32)
    i32.const 131783
    i32.const 13
    call 97)
  (func (;165;) (type 3)
    call 166
    unreachable)
  (func (;166;) (type 3)
    i32.const 131996
    i32.const 14
    call 0
    unreachable)
  (func (;167;) (type 3)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 30
    call 113
    call 112
    local.get 0
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 0
    i32.const 12
    i32.add
    i32.const 131146
    i32.const 26
    call 102
    local.get 0
    i32.load offset=4
    local.set 1
    local.get 0
    i32.load
    local.set 2
    local.get 0
    i32.load offset=12
    call 109
    local.get 2
    local.get 1
    call 81
    i32.const 1
    call 160
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;168;) (type 3)
    call 30
    i32.const 0
    call 111
    i32.const 1
    call 160)
  (func (;169;) (type 3)
    (local i32 i64 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i64 i32 i64)
    global.get 0
    i32.const 112
    i32.sub
    local.tee 0
    global.set 0
    i32.const 2
    call 111
    local.get 0
    i32.const 8
    i32.add
    call 103
    call 108
    local.set 1
    call 162
    call 93
    local.set 2
    i32.const -21
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        i32.const 0
        i32.load8_u offset=142028
        local.tee 4
        i32.eqz
        br_if 0 (;@2;)
        i32.const -21
        i32.const 2147483647
        local.get 4
        select
        local.set 3
        br 1 (;@1;)
      end
      i32.const 0
      i32.const 1
      i32.store8 offset=142028
      i32.const -21
      call 31
    end
    i32.const 1
    local.set 5
    block  ;; label = @1
      block  ;; label = @2
        local.get 3
        call 57
        i32.const 1
        i32.ne
        br_if 0 (;@2;)
        i32.const 0
        local.set 4
        local.get 3
        call 6
        local.set 6
        local.get 0
        i32.const 48
        i32.add
        local.set 7
        i32.const 0
        local.set 8
        loop  ;; label = @3
          local.get 8
          local.set 9
          local.get 4
          i32.const 16
          i32.add
          local.tee 10
          local.get 6
          i32.gt_u
          br_if 2 (;@1;)
          local.get 7
          i64.const 0
          i64.store
          local.get 0
          i64.const 0
          i64.store offset=40
          local.get 3
          local.get 4
          local.get 0
          i32.const 40
          i32.add
          i32.const 16
          call 101
          drop
          local.get 0
          i32.const 0
          i32.store offset=96
          i32.const 1
          local.set 8
          local.get 5
          i32.const 1
          i32.and
          local.set 11
          local.get 0
          i32.const 40
          i32.add
          local.get 0
          i32.const 96
          i32.add
          call 129
          local.set 12
          local.get 0
          i32.const 40
          i32.add
          local.get 0
          i32.const 96
          i32.add
          call 130
          local.set 13
          local.get 0
          i32.const 40
          i32.add
          local.get 0
          i32.const 96
          i32.add
          call 129
          local.set 14
          local.get 10
          local.set 4
          i32.const 0
          local.set 5
          local.get 11
          br_if 0 (;@3;)
        end
        local.get 9
        call 170
        unreachable
      end
      i32.const 131418
      i32.const 34
      call 0
      unreachable
    end
    block  ;; label = @1
      local.get 2
      call 83
      call 43
      call 39
      i32.eqz
      br_if 0 (;@1;)
      call 164
      call 145
      local.set 4
      call 164
      local.get 4
      i32.const 1
      i32.add
      local.tee 8
      i64.extend_i32_u
      local.tee 15
      call 158
      local.get 0
      i32.const 84
      i32.add
      call 66
      local.get 0
      i32.const 40
      i32.add
      local.get 0
      i32.load offset=84
      local.tee 10
      local.get 0
      i32.load offset=92
      local.tee 4
      local.get 8
      call 80
      local.get 4
      local.get 8
      call 77
      local.set 5
      call 143
      local.tee 4
      local.get 0
      i32.load offset=16
      call 2
      drop
      local.get 4
      local.get 0
      i32.load offset=20
      call 2
      drop
      local.get 0
      i32.load offset=24
      local.get 4
      call 92
      local.get 0
      i32.load offset=28
      local.get 4
      call 90
      local.get 0
      i64.load offset=8
      local.get 4
      call 64
      local.get 0
      i32.load offset=32
      local.get 4
      call 150
      local.get 5
      local.get 4
      call 144
      block  ;; label = @2
        local.get 10
        local.get 8
        call 141
        br_if 0 (;@2;)
        local.get 0
        i32.const 96
        i32.add
        local.get 0
        i32.load offset=88
        local.tee 5
        call 69
        local.get 0
        local.get 0
        i32.load offset=108
        i32.const 1
        i32.add
        local.tee 4
        i32.store offset=108
        block  ;; label = @3
          block  ;; label = @4
            local.get 0
            i32.load offset=96
            local.tee 3
            br_if 0 (;@4;)
            local.get 0
            local.get 4
            i32.store offset=100
            i32.const 0
            local.set 11
            br 1 (;@3;)
          end
          local.get 0
          local.get 5
          local.get 0
          i32.load offset=104
          local.tee 11
          call 68
          local.get 5
          local.get 11
          local.get 0
          i32.load
          local.get 4
          call 70
        end
        local.get 5
        local.get 4
        local.get 11
        i32.const 0
        call 70
        local.get 0
        local.get 4
        i32.store offset=104
        local.get 5
        i32.const 131680
        i32.const 6
        local.get 4
        call 71
        local.get 15
        call 158
        local.get 0
        local.get 3
        i32.const 1
        i32.add
        i32.store offset=96
        local.get 5
        local.get 0
        i32.const 96
        i32.add
        call 74
        local.get 10
        local.get 8
        call 75
        local.get 4
        i64.extend_i32_u
        call 158
      end
      local.get 8
      call 44
      local.set 5
      local.get 12
      call 143
      local.tee 4
      call 92
      local.get 13
      local.get 4
      call 64
      local.get 14
      local.get 4
      call 90
      local.get 5
      local.get 4
      call 144
      local.get 8
      call 60
      local.get 1
      call 158
      local.get 0
      i32.const 112
      i32.add
      global.set 0
      return
    end
    i32.const 131215
    i32.const 34
    call 59
    unreachable)
  (func (;170;) (type 5) (param i32)
    call 187
    unreachable)
  (func (;171;) (type 3)
    (local i32 i32 i32 i32 i32 i64 i64 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 256
    i32.sub
    local.tee 0
    global.set 0
    call 30
    i32.const 1
    call 111
    call 107
    local.set 1
    call 162
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          local.get 1
                          call 79
                          call 134
                          br_if 0 (;@11;)
                          local.get 0
                          i32.const 48
                          i32.add
                          local.get 1
                          call 41
                          local.get 0
                          i32.const 80
                          i32.add
                          local.get 1
                          call 44
                          call 45
                          local.get 0
                          i32.load offset=92
                          local.tee 2
                          call 56
                          br_if 1 (;@10;)
                          i32.const 2147483646
                          local.set 3
                          block  ;; label = @12
                            local.get 0
                            i32.load offset=72
                            local.tee 4
                            i32.const 2147483646
                            i32.ne
                            br_if 0 (;@12;)
                            i64.const 0
                            local.set 5
                            call 62
                            local.set 4
                            br 7 (;@5;)
                          end
                          local.get 4
                          call 6
                          local.set 3
                          local.get 0
                          i32.const 0
                          i32.store8 offset=216
                          local.get 0
                          local.get 3
                          i32.store offset=212
                          local.get 0
                          local.get 4
                          i32.store offset=208
                          local.get 0
                          local.get 3
                          i32.store offset=204
                          local.get 0
                          i32.const 0
                          i32.store offset=200
                          local.get 0
                          i32.const 96
                          i32.add
                          local.get 0
                          i32.const 200
                          i32.add
                          call 148
                          local.get 0
                          i32.load offset=96
                          br_if 5 (;@6;)
                          local.get 0
                          i32.load offset=100
                          local.set 4
                          local.get 0
                          i64.const 0
                          i64.store offset=96
                          local.get 0
                          i32.const 200
                          i32.add
                          local.get 0
                          i32.const 96
                          i32.add
                          i32.const 8
                          call 157
                          br_if 5 (;@6;)
                          local.get 0
                          i64.load offset=96
                          local.set 6
                          local.get 0
                          i32.const 0
                          i32.store8 offset=176
                          i32.const 1
                          local.set 7
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                local.get 0
                                i32.const 200
                                i32.add
                                local.get 0
                                i32.const 176
                                i32.add
                                i32.const 1
                                call 157
                                local.tee 3
                                br_if 0 (;@14;)
                                i32.const 4
                                local.set 8
                                local.get 0
                                i32.load8_u offset=176
                                local.tee 3
                                br_table 7 (;@7;) 2 (;@12;) 1 (;@13;)
                              end
                              local.get 0
                              local.get 3
                              i32.store offset=100
                              i32.const 15
                              local.set 3
                              br 5 (;@8;)
                            end
                            local.get 0
                            i32.const 131734
                            i32.store offset=100
                            i32.const 13
                            local.set 3
                            br 4 (;@8;)
                          end
                          local.get 0
                          i32.const 96
                          i32.add
                          local.get 0
                          i32.const 200
                          i32.add
                          call 149
                          local.get 0
                          i32.load offset=100
                          local.set 7
                          block  ;; label = @12
                            local.get 0
                            i32.load offset=96
                            local.tee 9
                            i32.eqz
                            br_if 0 (;@12;)
                            local.get 7
                            local.set 3
                            br 3 (;@9;)
                          end
                          call 62
                          local.set 3
                          block  ;; label = @12
                            loop  ;; label = @13
                              local.get 7
                              i32.eqz
                              br_if 1 (;@12;)
                              local.get 0
                              i32.const 96
                              i32.add
                              local.get 0
                              i32.const 200
                              i32.add
                              call 148
                              local.get 0
                              i32.load offset=100
                              local.set 8
                              block  ;; label = @14
                                local.get 0
                                i32.load offset=96
                                local.tee 9
                                i32.eqz
                                br_if 0 (;@14;)
                                local.get 8
                                local.set 3
                                br 5 (;@9;)
                              end
                              local.get 3
                              local.get 8
                              call 63
                              local.get 7
                              i32.const -1
                              i32.add
                              local.set 7
                              br 0 (;@13;)
                            end
                          end
                          i32.const 1
                          local.set 7
                          local.get 0
                          i32.const 1
                          i32.store offset=100
                          i32.const 8
                          local.set 8
                          br 4 (;@7;)
                        end
                        i32.const 131249
                        i32.const 37
                        call 59
                        unreachable
                      end
                      i32.const 131286
                      i32.const 17
                      call 59
                      unreachable
                    end
                    local.get 0
                    local.get 9
                    i32.store offset=100
                  end
                  i32.const 8
                  local.set 8
                  i32.const 0
                  local.set 7
                end
                local.get 0
                i32.const 96
                i32.add
                local.get 8
                i32.add
                local.get 3
                i32.store
                local.get 0
                i32.load offset=104
                local.set 3
                local.get 0
                i32.load offset=100
                local.set 8
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 7
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 3
                    i32.const 2147483646
                    local.get 8
                    i32.const 1
                    i32.eq
                    select
                    local.set 3
                    br 1 (;@7;)
                  end
                  local.get 8
                  br_if 1 (;@6;)
                end
                local.get 0
                i32.load offset=204
                local.get 0
                i32.load offset=200
                i32.ne
                br_if 0 (;@6;)
                local.get 6
                i64.const 8
                i64.shr_u
                i64.const 4278190080
                i64.and
                local.get 6
                i64.const 24
                i64.shr_u
                i64.const 16711680
                i64.and
                i64.or
                local.get 6
                i64.const 40
                i64.shr_u
                i64.const 65280
                i64.and
                local.get 6
                i64.const 56
                i64.shr_u
                i64.or
                i64.or
                local.set 5
                local.get 6
                i64.const 56
                i64.shl
                local.get 6
                i64.const 65280
                i64.and
                i64.const 40
                i64.shl
                i64.or
                local.get 6
                i64.const 16711680
                i64.and
                i64.const 24
                i64.shl
                local.get 6
                i64.const 4278190080
                i64.and
                i64.const 8
                i64.shl
                i64.or
                i64.or
                local.set 6
                block  ;; label = @7
                  local.get 0
                  i32.load8_u offset=216
                  i32.eqz
                  br_if 0 (;@7;)
                  i32.const 0
                  i32.const 0
                  i32.store offset=142016
                  i32.const 0
                  i32.const 0
                  i32.store8 offset=142020
                end
                local.get 6
                local.get 5
                i64.or
                local.set 5
                br 1 (;@5;)
              end
              local.get 0
              i32.load8_u offset=216
              i32.eqz
              br_if 1 (;@4;)
              i32.const 0
              i32.const 0
              i32.store offset=142016
              i32.const 0
              i32.const 0
              i32.store8 offset=142020
              br 1 (;@4;)
            end
            local.get 4
            call 118
            br_if 0 (;@4;)
            local.get 5
            i64.const -250000000
            i64.add
            i64.const -240000000
            i64.ge_u
            br_if 1 (;@3;)
          end
          local.get 1
          call 40
          local.get 1
          call 65
          br 1 (;@2;)
        end
        call 28
        local.get 5
        i64.const 20000000
        i64.add
        i64.le_u
        br_if 1 (;@1;)
        call 62
        local.set 10
        call 62
        local.tee 8
        local.get 1
        call 163
        call 62
        local.set 9
        local.get 0
        i64.load offset=80
        local.set 6
        local.get 0
        i32.load offset=88
        local.set 7
        local.get 0
        local.get 2
        i32.const 24
        i32.shl
        local.get 2
        i32.const 65280
        i32.and
        i32.const 8
        i32.shl
        i32.or
        local.get 2
        i32.const 8
        i32.shr_u
        i32.const 65280
        i32.and
        local.get 2
        i32.const 24
        i32.shr_u
        i32.or
        i32.or
        i32.store offset=212
        local.get 0
        local.get 7
        i32.const 24
        i32.shl
        local.get 7
        i32.const 65280
        i32.and
        i32.const 8
        i32.shl
        i32.or
        local.get 7
        i32.const 8
        i32.shr_u
        i32.const 65280
        i32.and
        local.get 7
        i32.const 24
        i32.shr_u
        i32.or
        i32.or
        i32.store offset=200
        local.get 0
        local.get 6
        i64.const 56
        i64.shl
        local.get 6
        i64.const 65280
        i64.and
        i64.const 40
        i64.shl
        i64.or
        local.get 6
        i64.const 16711680
        i64.and
        i64.const 24
        i64.shl
        local.get 6
        i64.const 4278190080
        i64.and
        i64.const 8
        i64.shl
        i64.or
        i64.or
        local.get 6
        i64.const 8
        i64.shr_u
        i64.const 4278190080
        i64.and
        local.get 6
        i64.const 24
        i64.shr_u
        i64.const 16711680
        i64.and
        i64.or
        local.get 6
        i64.const 40
        i64.shr_u
        i64.const 65280
        i64.and
        local.get 6
        i64.const 56
        i64.shr_u
        i64.or
        i64.or
        i64.or
        i64.store offset=204 align=4
        local.get 9
        local.get 0
        i32.const 200
        i32.add
        i32.const 16
        call 12
        drop
        call 32
        local.set 6
        local.get 1
        call 79
        local.get 6
        call 158
        local.get 0
        local.get 0
        i32.const 60
        i32.add
        i32.store offset=140
        local.get 0
        local.get 9
        i32.store offset=136
        local.get 0
        local.get 8
        i32.store offset=128
        local.get 0
        i32.const 18
        i32.store offset=124
        local.get 0
        i32.const 131086
        i32.store offset=120
        local.get 0
        i64.const 20000000
        i64.store offset=112
        local.get 0
        local.get 10
        local.get 3
        local.get 3
        i32.const 2147483646
        i32.eq
        select
        i32.store offset=108
        local.get 0
        local.get 4
        i32.store offset=104
        local.get 0
        local.get 5
        i64.store offset=96
        i32.const -25
        i32.const 131960
        i32.const 0
        call 20
        drop
        i32.const -25
        call 38
        local.set 1
        local.get 0
        local.get 8
        call 6
        i32.store offset=208
        local.get 0
        i32.const 0
        i32.store offset=204
        local.get 0
        local.get 0
        i32.const 128
        i32.add
        i32.store offset=200
        block  ;; label = @3
          loop  ;; label = @4
            local.get 0
            i32.const 40
            i32.add
            local.get 0
            i32.const 200
            i32.add
            call 152
            local.get 0
            i32.load offset=40
            i32.eqz
            br_if 1 (;@3;)
            local.get 0
            i32.load offset=44
            local.get 1
            call 92
            br 0 (;@4;)
          end
        end
        local.get 0
        local.get 1
        i32.store offset=148
        local.get 0
        local.get 0
        i64.load offset=112
        i64.store offset=152
        local.get 0
        local.get 0
        i64.load offset=96
        i64.store offset=160
        local.get 0
        i32.load offset=140
        local.set 3
        local.get 0
        i32.load offset=104
        local.set 10
        local.get 0
        i32.load offset=108
        local.set 11
        local.get 0
        local.get 0
        i32.load offset=136
        local.tee 1
        i32.store offset=172
        local.get 0
        i32.const 144
        i32.add
        local.set 2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 1
              call 57
              br_table 1 (;@4;) 2 (;@3;) 0 (;@5;)
            end
            local.get 0
            local.get 11
            i32.store offset=236
            local.get 0
            local.get 10
            i32.store offset=232
            local.get 0
            local.get 2
            i32.store offset=228
            local.get 0
            local.get 2
            i32.store offset=224
            local.get 0
            i32.const 18
            i32.store offset=208
            local.get 0
            i32.const 131086
            i32.store offset=204
            local.get 0
            local.get 0
            i32.const 148
            i32.add
            i32.store offset=220
            local.get 0
            local.get 0
            i32.const 152
            i32.add
            i32.store offset=216
            local.get 0
            local.get 0
            i32.const 160
            i32.add
            i32.store offset=212
            local.get 0
            local.get 0
            i32.const 172
            i32.add
            i32.store offset=200
            local.get 0
            i32.const 16
            i32.add
            i32.const 131563
            i32.const 20
            call 47
            local.get 0
            i32.load offset=16
            local.set 4
            local.get 0
            i32.load offset=20
            local.set 1
            call 62
            drop
            local.get 1
            local.get 3
            i32.load
            call 38
            call 63
            local.get 1
            local.get 0
            i32.load offset=172
            call 57
            call 163
            local.get 0
            i32.const 204
            i32.add
            local.set 12
            local.get 0
            i32.load offset=172
            call 6
            local.set 8
            local.get 0
            i32.const 184
            i32.add
            local.set 9
            i32.const 0
            local.set 2
            block  ;; label = @5
              loop  ;; label = @6
                local.get 2
                i32.const 16
                i32.add
                local.tee 3
                local.get 8
                i32.gt_u
                br_if 1 (;@5;)
                local.get 9
                i64.const 0
                i64.store
                local.get 0
                i64.const 0
                i64.store offset=176
                local.get 0
                i32.load offset=172
                local.get 2
                local.get 0
                i32.const 176
                i32.add
                i32.const 16
                call 101
                drop
                local.get 0
                i32.const 0
                i32.store offset=196
                local.get 0
                i32.const 176
                i32.add
                local.get 0
                i32.const 196
                i32.add
                call 129
                local.set 2
                local.get 0
                i32.const 176
                i32.add
                local.get 0
                i32.const 196
                i32.add
                call 130
                local.set 6
                local.get 0
                i32.const 176
                i32.add
                local.get 0
                i32.const 196
                i32.add
                call 129
                local.set 7
                call 62
                drop
                local.get 1
                local.get 2
                call 38
                call 63
                local.get 0
                i32.const 8
                i32.add
                local.get 4
                local.get 1
                local.get 6
                call 115
                local.get 0
                i32.load offset=8
                local.set 4
                local.get 0
                i32.load offset=12
                local.set 1
                call 62
                drop
                local.get 1
                local.get 7
                call 91
                call 63
                local.get 3
                local.set 2
                br 0 (;@6;)
              end
            end
            local.get 0
            local.get 4
            local.get 1
            local.get 10
            local.get 11
            call 117
            local.get 0
            i32.load offset=4
            local.set 1
            local.get 0
            i32.load
            local.set 2
            local.get 12
            call 54
            call 51
            local.get 2
            local.get 1
            call 122
            br 2 (;@2;)
          end
          local.get 0
          local.get 11
          i32.store offset=224
          local.get 0
          local.get 10
          i32.store offset=220
          local.get 0
          i32.const 18
          i32.store offset=204
          local.get 0
          i32.const 131086
          i32.store offset=200
          local.get 0
          local.get 0
          i32.const 148
          i32.add
          i32.store offset=216
          local.get 0
          local.get 0
          i32.const 152
          i32.add
          i32.store offset=212
          local.get 0
          local.get 0
          i32.const 160
          i32.add
          i32.store offset=208
          call 51
          local.set 1
          local.get 0
          i32.const 200
          i32.add
          local.get 3
          i32.load
          local.get 1
          local.get 10
          local.get 11
          call 122
          br 1 (;@2;)
        end
        local.get 0
        i32.const 176
        i32.add
        local.get 1
        call 58
        local.get 0
        i64.load offset=176
        local.set 6
        local.get 0
        local.get 0
        i32.const 188
        i32.add
        local.tee 4
        i32.store offset=220
        local.get 0
        local.get 0
        i32.const 176
        i32.add
        i32.const 8
        i32.add
        local.tee 7
        i32.store offset=216
        local.get 0
        local.get 6
        i64.store offset=208
        local.get 0
        local.get 2
        i32.store offset=248
        local.get 0
        local.get 2
        i32.store offset=244
        local.get 0
        i32.const 18
        i32.store offset=228
        local.get 0
        i32.const 131086
        i32.store offset=224
        local.get 0
        local.get 11
        i32.store offset=204
        local.get 0
        local.get 10
        i32.store offset=200
        local.get 0
        local.get 0
        i32.const 148
        i32.add
        i32.store offset=240
        local.get 0
        local.get 0
        i32.const 152
        i32.add
        i32.store offset=236
        local.get 0
        local.get 0
        i32.const 160
        i32.add
        i32.store offset=232
        local.get 0
        i32.const 224
        i32.add
        local.set 1
        block  ;; label = @3
          local.get 6
          i64.const 0
          i64.ne
          br_if 0 (;@3;)
          local.get 0
          i32.const 24
          i32.add
          local.get 10
          local.get 11
          local.get 7
          local.get 4
          call 50
          local.get 0
          i32.load offset=28
          local.set 2
          local.get 0
          i32.load offset=24
          local.set 4
          call 51
          local.set 7
          local.get 1
          local.get 3
          i32.load
          local.get 7
          local.get 4
          local.get 2
          call 122
          br 1 (;@2;)
        end
        local.get 0
        i32.const 32
        i32.add
        local.get 10
        local.get 11
        local.get 3
        i32.load
        local.get 0
        i32.const 200
        i32.add
        i32.const 8
        i32.add
        call 53
        local.get 0
        i32.load offset=36
        local.set 2
        local.get 0
        i32.load offset=32
        local.set 4
        local.get 1
        call 54
        call 51
        local.get 4
        local.get 2
        call 122
      end
      local.get 0
      i32.const 256
      i32.add
      global.set 0
      return
    end
    i32.const 131303
    i32.const 25
    call 59
    unreachable)
  (func (;172;) (type 3)
    (local i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    call 30
    i32.const 1
    call 111
    local.get 0
    call 107
    call 41
    i32.const 131960
    i32.const 0
    call 97
    local.tee 1
    local.get 0
    i32.load offset=8
    call 2
    drop
    local.get 1
    local.get 0
    i32.load offset=12
    call 2
    drop
    local.get 0
    i32.load offset=16
    local.get 1
    call 92
    local.get 0
    i32.load offset=20
    local.get 1
    call 90
    local.get 0
    i64.load
    local.get 1
    call 64
    local.get 0
    i32.load offset=24
    local.get 1
    call 150
    local.get 1
    call 16
    drop
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;173;) (type 3)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i64 i32 i32 i32 i32)
    global.get 0
    i32.const 80
    i32.sub
    local.tee 0
    global.set 0
    call 30
    i32.const 0
    call 111
    call 62
    local.set 1
    local.get 0
    i32.const 28
    i32.add
    call 66
    local.get 0
    i32.const 40
    i32.add
    local.get 0
    i32.load offset=32
    local.tee 2
    call 69
    local.get 0
    i32.load offset=44
    local.set 3
    local.get 0
    i32.load offset=36
    local.set 4
    local.get 0
    i32.load offset=28
    local.set 5
    block  ;; label = @1
      block  ;; label = @2
        loop  ;; label = @3
          local.get 3
          i32.eqz
          br_if 1 (;@2;)
          local.get 0
          i32.const 16
          i32.add
          local.get 2
          local.get 3
          call 68
          local.get 0
          i32.load offset=20
          local.set 6
          local.get 0
          i32.const 40
          i32.add
          local.get 5
          local.get 4
          local.get 2
          local.get 3
          call 73
          local.tee 3
          call 80
          local.get 0
          i64.load offset=40
          i64.const 0
          i64.eq
          br_if 2 (;@1;)
          local.get 0
          i32.load offset=72
          local.set 7
          local.get 0
          i64.load offset=48
          local.set 8
          local.get 0
          i32.load offset=68
          local.set 9
          local.get 0
          i32.load offset=64
          local.set 10
          local.get 0
          i32.load offset=60
          local.set 11
          local.get 0
          i32.load offset=56
          local.set 12
          local.get 1
          local.get 3
          call 163
          call 62
          call 38
          local.tee 3
          local.get 12
          call 2
          drop
          local.get 3
          local.get 11
          call 2
          drop
          local.get 10
          local.get 3
          call 92
          local.get 9
          call 91
          local.get 3
          call 92
          local.get 8
          local.get 3
          call 64
          block  ;; label = @4
            block  ;; label = @5
              local.get 7
              i32.const 2147483646
              i32.ne
              br_if 0 (;@5;)
              local.get 3
              i32.const 0
              call 151
              br 1 (;@4;)
            end
            local.get 3
            i32.const 1
            call 151
            local.get 7
            local.get 3
            call 92
          end
          local.get 1
          local.get 3
          call 63
          local.get 6
          local.set 3
          br 0 (;@3;)
        end
      end
      local.get 0
      local.get 1
      i32.store offset=24
      local.get 0
      local.get 1
      call 6
      i32.store offset=48
      local.get 0
      i32.const 0
      i32.store offset=44
      local.get 0
      local.get 0
      i32.const 24
      i32.add
      i32.store offset=40
      block  ;; label = @2
        loop  ;; label = @3
          local.get 0
          i32.const 8
          i32.add
          local.get 0
          i32.const 40
          i32.add
          call 152
          local.get 0
          i32.load offset=8
          i32.eqz
          br_if 1 (;@2;)
          local.get 0
          i32.load offset=12
          call 16
          drop
          br 0 (;@3;)
        end
      end
      local.get 0
      i32.const 80
      i32.add
      global.set 0
      return
    end
    call 174
    unreachable)
  (func (;174;) (type 3)
    call 188
    unreachable)
  (func (;175;) (type 3)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 30
    call 94
    call 113
    call 112
    local.get 0
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 0
    i32.const 12
    i32.add
    i32.const 131146
    i32.const 26
    call 102
    local.get 0
    i32.load offset=4
    local.set 1
    local.get 0
    i32.load
    local.set 2
    local.get 0
    i32.load offset=12
    call 109
    local.get 2
    local.get 1
    call 81
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;176;) (type 3)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 30
    call 94
    call 113
    call 112
    local.get 0
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 0
    i32.const 12
    i32.add
    i32.const 131407
    i32.const 11
    call 102
    local.get 0
    i32.load offset=4
    local.set 1
    local.get 0
    i32.load
    local.set 2
    local.get 0
    i32.load offset=12
    call 109
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          br_if 0 (;@3;)
          local.get 1
          call 82
          i32.eqz
          br_if 1 (;@2;)
          call 46
          local.get 1
          call 84
          br 2 (;@1;)
        end
        call 46
        call 78
        br 1 (;@1;)
      end
      i32.const 131339
      i32.const 38
      call 59
      unreachable
    end
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;177;) (type 3)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 30
    call 94
    call 113
    call 112
    local.get 0
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 0
    i32.const 12
    i32.add
    i32.const 131407
    i32.const 11
    call 102
    local.get 0
    i32.load offset=4
    local.set 1
    local.get 0
    i32.load
    local.set 2
    local.get 0
    i32.load offset=12
    call 109
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          br_if 0 (;@3;)
          local.get 1
          call 82
          i32.eqz
          br_if 1 (;@2;)
          call 42
          local.get 1
          call 84
          br 2 (;@1;)
        end
        call 42
        call 78
        br 1 (;@1;)
      end
      i32.const 131339
      i32.const 38
      call 59
      unreachable
    end
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;178;) (type 3)
    call 30
    i32.const 0
    call 111
    call 83
    call 114)
  (func (;179;) (type 3)
    call 30
    i32.const 0
    call 111
    call 46
    call 114)
  (func (;180;) (type 3)
    call 30
    i32.const 0
    call 111
    call 42
    call 114)
  (func (;181;) (type 3)
    call 30
    i32.const 0
    call 111
    call 164
    call 145
    i64.extend_i32_u
    call 33)
  (func (;182;) (type 3)
    call 30
    call 94
    i32.const 0
    call 111
    i32.const 1
    call 160
    i32.const 131910
    i32.const 13
    call 147
    call 62
    call 34)
  (func (;183;) (type 3)
    call 30
    call 94
    i32.const 0
    call 111
    i32.const 0
    call 160
    i32.const 131942
    i32.const 15
    call 147
    call 62
    call 34)
  (func (;184;) (type 3)
    call 30
    i32.const 0
    call 111
    call 161
    call 146
    i64.extend_i32_u
    call 35)
  (func (;185;) (type 3)
    (local i32 i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 0
    global.set 0
    i32.const -25
    call 36
    i32.const -23
    i32.const 131960
    i32.const 0
    call 20
    drop
    i32.const -25
    call 6
    local.set 1
    local.get 0
    i32.const 0
    i32.store8 offset=44
    local.get 0
    local.get 1
    i32.store offset=40
    local.get 0
    i32.const -25
    i32.store offset=36
    local.get 0
    local.get 1
    i32.store offset=32
    local.get 0
    i32.const 0
    i32.store offset=28
    local.get 0
    i32.const 28
    i32.add
    i32.const 8
    i32.add
    local.set 2
    i32.const 0
    local.set 3
    block  ;; label = @1
      loop  ;; label = @2
        block  ;; label = @3
          local.get 1
          local.get 3
          i32.ne
          br_if 0 (;@3;)
          block  ;; label = @4
            local.get 0
            i32.load8_u offset=44
            i32.eqz
            br_if 0 (;@4;)
            i32.const 0
            i32.const 0
            i32.store offset=142016
            i32.const 0
            i32.const 0
            i32.store8 offset=142020
          end
          i32.const -23
          call 6
          local.set 3
          local.get 0
          i32.const 0
          i32.store offset=36
          local.get 0
          i32.const -23
          i32.store offset=28
          local.get 0
          local.get 3
          i32.const 2
          i32.shr_u
          i32.store offset=32
          local.get 0
          i32.const 28
          i32.add
          call 100
          local.set 3
          local.get 0
          i32.load offset=32
          local.get 0
          i32.load offset=36
          call 110
          call 113
          call 112
          local.get 0
          i32.const 0
          i32.store offset=12
          local.get 0
          i32.const 16
          i32.add
          local.get 0
          i32.const 12
          i32.add
          call 98
          local.get 0
          i32.load offset=12
          call 109
          block  ;; label = @4
            local.get 0
            i32.load offset=16
            i32.eqz
            br_if 0 (;@4;)
            local.get 3
            call 40
          end
          local.get 3
          call 65
          local.get 0
          i32.const 48
          i32.add
          global.set 0
          return
        end
        local.get 0
        i32.const 0
        i32.store offset=16
        block  ;; label = @3
          local.get 2
          local.get 3
          local.get 0
          i32.const 16
          i32.add
          i32.const 4
          call 132
          br_if 0 (;@3;)
          local.get 0
          local.get 0
          i32.load offset=36
          local.get 0
          i32.load offset=28
          i32.const 4
          i32.add
          local.tee 1
          local.get 0
          i32.load offset=16
          local.tee 3
          i32.const 24
          i32.shl
          local.get 3
          i32.const 65280
          i32.and
          i32.const 8
          i32.shl
          i32.or
          local.get 3
          i32.const 8
          i32.shr_u
          i32.const 65280
          i32.and
          local.get 3
          i32.const 24
          i32.shr_u
          i32.or
          i32.or
          local.tee 3
          call 125
          local.get 0
          i32.load
          i32.const 1
          i32.ne
          br_if 2 (;@1;)
          local.get 0
          i32.load offset=4
          local.set 4
          local.get 0
          local.get 3
          local.get 1
          i32.add
          local.tee 3
          i32.store offset=28
          i32.const -23
          local.get 4
          call 63
          local.get 0
          i32.load offset=32
          local.set 1
          br 1 (;@2;)
        end
      end
      i32.const 131538
      i32.const 25
      call 156
      unreachable
    end
    i32.const 131538
    i32.const 25
    call 153
    unreachable)
  (func (;186;) (type 3))
  (func (;187;) (type 3)
    call 165
    unreachable)
  (func (;188;) (type 3)
    call 187
    unreachable)
  (func (;189;) (type 0) (param i32 i32)
    call 187
    unreachable)
  (func (;190;) (type 0) (param i32 i32)
    call 187
    unreachable)
  (func (;191;) (type 6) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 16
        i32.ge_u
        br_if 0 (;@2;)
        local.get 0
        local.set 3
        br 1 (;@1;)
      end
      local.get 0
      i32.const 0
      local.get 0
      i32.sub
      i32.const 3
      i32.and
      local.tee 4
      i32.add
      local.set 5
      block  ;; label = @2
        local.get 4
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        local.set 3
        local.get 1
        local.set 6
        loop  ;; label = @3
          local.get 3
          local.get 6
          i32.load8_u
          i32.store8
          local.get 6
          i32.const 1
          i32.add
          local.set 6
          local.get 3
          i32.const 1
          i32.add
          local.tee 3
          local.get 5
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 5
      local.get 2
      local.get 4
      i32.sub
      local.tee 7
      i32.const -4
      i32.and
      local.tee 8
      i32.add
      local.set 3
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          local.get 4
          i32.add
          local.tee 9
          i32.const 3
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 8
          i32.const 1
          i32.lt_s
          br_if 1 (;@2;)
          local.get 9
          i32.const 3
          i32.shl
          local.tee 6
          i32.const 24
          i32.and
          local.set 2
          local.get 9
          i32.const -4
          i32.and
          local.tee 10
          i32.const 4
          i32.add
          local.set 1
          i32.const 0
          local.get 6
          i32.sub
          i32.const 24
          i32.and
          local.set 4
          local.get 10
          i32.load
          local.set 6
          loop  ;; label = @4
            local.get 5
            local.get 6
            local.get 2
            i32.shr_u
            local.get 1
            i32.load
            local.tee 6
            local.get 4
            i32.shl
            i32.or
            i32.store
            local.get 1
            i32.const 4
            i32.add
            local.set 1
            local.get 5
            i32.const 4
            i32.add
            local.tee 5
            local.get 3
            i32.lt_u
            br_if 0 (;@4;)
            br 2 (;@2;)
          end
        end
        local.get 8
        i32.const 1
        i32.lt_s
        br_if 0 (;@2;)
        local.get 9
        local.set 1
        loop  ;; label = @3
          local.get 5
          local.get 1
          i32.load
          i32.store
          local.get 1
          i32.const 4
          i32.add
          local.set 1
          local.get 5
          i32.const 4
          i32.add
          local.tee 5
          local.get 3
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 7
      i32.const 3
      i32.and
      local.set 2
      local.get 9
      local.get 8
      i32.add
      local.set 1
    end
    block  ;; label = @1
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 3
      local.get 2
      i32.add
      local.set 5
      loop  ;; label = @2
        local.get 3
        local.get 1
        i32.load8_u
        i32.store8
        local.get 1
        i32.const 1
        i32.add
        local.set 1
        local.get 3
        i32.const 1
        i32.add
        local.tee 3
        local.get 5
        i32.lt_u
        br_if 0 (;@2;)
      end
    end
    local.get 0)
  (func (;192;) (type 6) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call 191)
  (memory (;0;) 3)
  (global (;0;) (mut i32) (i32.const 131072))
  (global (;1;) i32 (i32.const 142061))
  (global (;2;) i32 (i32.const 142064))
  (export "memory" (memory 0))
  (export "init" (func 167))
  (export "upgrade" (func 168))
  (export "deposit" (func 169))
  (export "execute" (func 171))
  (export "getPendingTransactionById" (func 172))
  (export "getPendingTransactions" (func 173))
  (export "setMultiTransferAddress" (func 175))
  (export "setBridgedTokensWrapperAddress" (func 176))
  (export "setDcdtSafeAddress" (func 177))
  (export "getMultiTransferAddress" (func 178))
  (export "getBridgedTokensWrapperAddress" (func 179))
  (export "getDcdtSafeContractAddress" (func 180))
  (export "highestTxId" (func 181))
  (export "pause" (func 182))
  (export "unpause" (func 183))
  (export "isPaused" (func 184))
  (export "execution_callback" (func 185))
  (export "callBack" (func 186))
  (export "__data_end" (global 1))
  (export "__heap_base" (global 2))
  (data (;0;) (i32.const 131072) "input too longexecution_callbackcreateTransactionbatch_ideth_txtx_idresultopt_multi_transfer_addressExpected only one dcdt paymentInvalid tx idOnly MultiTransfer can do depositsTransaction is already being executedNo amount bridgedNot enough gas to executeunwrapTokenInvalid bridged tokens wrapper addressInvalid multi-transfer addressopt_addressincorrect number of DCDT transfersargument decode error (): too few argumentstoo many argumentswrong number of argumentsserializer decode error: MultiDCDTNFTTransferDCDTNFTTransferDCDTTransferinput too shortManagedVec index out of range.mapped.node_id.node_links.value.infostorage decode error (key: bad array lengthinvalid valueinput out of rangeContract is pausedhighest_tx_idongoingExecutionpending_transactionsmultiTransferAddressdcdtSafeContractAddressbridgedTokensWrapperAddresspaymentspauseContractpause_module:pausedunpauseContract\00\00\00Endpoint can only be called by ownerpanic occurred")
  (data (;1;) (i32.const 132012) "8\ff\ff\ff"))
