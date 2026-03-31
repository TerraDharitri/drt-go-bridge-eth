(module
  (type (;0;) (func (param i32 i64)))
  (type (;1;) (func (param i32 i32 i32)))
  (type (;2;) (func (param i32 i32)))
  (type (;3;) (func (result i32)))
  (type (;4;) (func (param i32 i32) (result i32)))
  (type (;5;) (func (param i64 i32 i32 i32 i32 i32) (result i32)))
  (type (;6;) (func (param i32 i32 i64 i32 i32) (result i32)))
  (type (;7;) (func (param i32)))
  (type (;8;) (func (param i32) (result i32)))
  (type (;9;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;10;) (func (param i32 i32 i32) (result i32)))
  (type (;11;) (func (param i32) (result i64)))
  (type (;12;) (func (result i64)))
  (type (;13;) (func))
  (type (;14;) (func (param i64)))
  (type (;15;) (func (param i64 i32 i32 i32 i32)))
  (type (;16;) (func (param i32 i32 i32 i32)))
  (type (;17;) (func (param i32 i32) (result i64)))
  (type (;18;) (func (param i64) (result i32)))
  (type (;19;) (func (param i64 i32)))
  (type (;20;) (func (param i32 i64 i32 i32)))
  (type (;21;) (func (param i64 i32) (result i32)))
  (import "env" "bigIntSetInt64" (func (;0;) (type 0)))
  (import "env" "bigIntAdd" (func (;1;) (type 1)))
  (import "env" "signalError" (func (;2;) (type 2)))
  (import "env" "mBufferNew" (func (;3;) (type 3)))
  (import "env" "mBufferAppend" (func (;4;) (type 4)))
  (import "env" "mBufferEq" (func (;5;) (type 4)))
  (import "env" "managedExecuteOnDestContext" (func (;6;) (type 5)))
  (import "env" "managedMultiTransferDCDTNFTExecute" (func (;7;) (type 6)))
  (import "env" "managedCaller" (func (;8;) (type 7)))
  (import "env" "managedOwnerAddress" (func (;9;) (type 7)))
  (import "env" "managedGetMultiDCDTCallValue" (func (;10;) (type 7)))
  (import "env" "mBufferGetLength" (func (;11;) (type 8)))
  (import "env" "mBufferGetByteSlice" (func (;12;) (type 9)))
  (import "env" "mBufferGetArgument" (func (;13;) (type 4)))
  (import "env" "mBufferAppendBytes" (func (;14;) (type 10)))
  (import "env" "managedSignalError" (func (;15;) (type 7)))
  (import "env" "mBufferCopyByteSlice" (func (;16;) (type 9)))
  (import "env" "smallIntGetUnsignedArgument" (func (;17;) (type 11)))
  (import "env" "getNumArguments" (func (;18;) (type 3)))
  (import "env" "mBufferFinish" (func (;19;) (type 8)))
  (import "env" "getGasLeft" (func (;20;) (type 12)))
  (import "env" "managedSCAddress" (func (;21;) (type 7)))
  (import "env" "cleanReturnData" (func (;22;) (type 13)))
  (import "env" "mBufferSetBytes" (func (;23;) (type 10)))
  (import "env" "mBufferFromBigIntUnsigned" (func (;24;) (type 4)))
  (import "env" "mBufferStorageLoad" (func (;25;) (type 4)))
  (import "env" "mBufferStorageStore" (func (;26;) (type 4)))
  (import "env" "mBufferToBigIntUnsigned" (func (;27;) (type 4)))
  (import "env" "bigIntCmp" (func (;28;) (type 4)))
  (import "env" "bigIntSub" (func (;29;) (type 1)))
  (import "env" "managedWriteLog" (func (;30;) (type 2)))
  (import "env" "bigIntMul" (func (;31;) (type 1)))
  (import "env" "bigIntTDiv" (func (;32;) (type 1)))
  (import "env" "getDCDTLocalRoles" (func (;33;) (type 11)))
  (import "env" "checkNoPayment" (func (;34;) (type 13)))
  (import "env" "bigIntFinishUnsigned" (func (;35;) (type 7)))
  (import "env" "smallIntFinishSigned" (func (;36;) (type 14)))
  (import "env" "bigIntSign" (func (;37;) (type 8)))
  (func (;38;) (type 8) (param i32) (result i32)
    (local i32)
    call 39
    local.tee 1
    i64.const 0
    call 0
    local.get 1
    local.get 1
    local.get 0
    call 1
    local.get 1)
  (func (;39;) (type 3) (result i32)
    (local i32)
    i32.const 0
    i32.const 0
    i32.load offset=132128
    i32.const -1
    i32.add
    local.tee 0
    i32.store offset=132128
    local.get 0)
  (func (;40;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 2
    unreachable)
  (func (;41;) (type 8) (param i32) (result i32)
    (local i32)
    call 3
    local.tee 1
    local.get 0
    call 4
    drop
    local.get 1)
  (func (;42;) (type 8) (param i32) (result i32)
    (local i32 i32 i64 i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 1
    global.set 0
    call 43
    local.set 2
    local.get 1
    local.get 0
    i32.load
    call 44
    i32.store offset=20
    local.get 1
    i32.const 0
    i32.store offset=16
    local.get 1
    local.get 0
    i32.store offset=12
    loop (result i32)  ;; label = @1
      local.get 1
      i32.const 24
      i32.add
      local.get 1
      i32.const 12
      i32.add
      call 45
      block  ;; label = @2
        local.get 1
        i32.load offset=24
        br_if 0 (;@2;)
        local.get 1
        i32.const 64
        i32.add
        global.set 0
        local.get 2
        return
      end
      local.get 1
      i64.load offset=32
      local.set 3
      local.get 1
      i32.load offset=44
      local.set 0
      local.get 1
      i32.load offset=40
      call 41
      local.set 4
      local.get 1
      local.get 0
      call 38
      i32.store offset=60
      local.get 1
      local.get 3
      i64.store offset=48
      local.get 1
      local.get 4
      i32.store offset=56
      local.get 2
      local.get 1
      i32.const 48
      i32.add
      call 46
      br 0 (;@1;)
    end)
  (func (;43;) (type 3) (result i32)
    (local i32)
    call 39
    local.tee 0
    i32.const 1
    i32.const 0
    call 23
    drop
    local.get 0)
  (func (;44;) (type 8) (param i32) (result i32)
    local.get 0
    call 11)
  (func (;45;) (type 2) (param i32 i32)
    (local i32 i64 i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    i64.const 0
    local.set 3
    block  ;; label = @1
      local.get 1
      i32.load offset=4
      local.tee 4
      i32.const 16
      i32.add
      local.tee 5
      local.get 1
      i32.load offset=8
      i32.gt_u
      br_if 0 (;@1;)
      local.get 1
      i32.load
      local.set 6
      local.get 2
      i32.const 16
      i32.add
      i64.const 0
      i64.store
      local.get 2
      i64.const 0
      i64.store offset=8
      local.get 6
      i32.load
      local.get 4
      local.get 2
      i32.const 8
      i32.add
      i32.const 16
      call 78
      local.get 2
      i32.const 0
      i32.store offset=28
      local.get 2
      i32.const 8
      i32.add
      local.get 2
      i32.const 28
      i32.add
      call 57
      local.set 4
      local.get 2
      i32.const 8
      i32.add
      local.get 2
      i32.const 28
      i32.add
      call 58
      local.set 3
      local.get 0
      local.get 2
      i32.const 8
      i32.add
      local.get 2
      i32.const 28
      i32.add
      call 57
      i32.store offset=20
      local.get 0
      local.get 4
      i32.store offset=16
      local.get 0
      local.get 3
      i64.store offset=8
      local.get 1
      local.get 5
      i32.store offset=4
      i64.const 1
      local.set 3
    end
    local.get 0
    local.get 3
    i64.store
    local.get 2
    i32.const 32
    i32.add
    global.set 0)
  (func (;46;) (type 2) (param i32 i32)
    (local i32 i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.load offset=8
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
    i32.store
    local.get 2
    local.get 1
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
    i32.store offset=12
    local.get 2
    local.get 1
    i64.load
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
    i64.store offset=4 align=4
    local.get 0
    local.get 2
    i32.const 16
    call 14
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;47;) (type 4) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 5
    i32.const 0
    i32.gt_s)
  (func (;48;) (type 2) (param i32 i32)
    (local i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.load offset=4
        local.tee 2
        local.get 1
        i32.load offset=8
        i32.le_u
        br_if 0 (;@2;)
        i32.const 0
        local.set 3
        br 1 (;@1;)
      end
      i32.const 1
      local.set 3
      local.get 1
      local.get 2
      i32.const 1
      i32.add
      i32.store offset=4
      local.get 1
      i32.load
      i32.load
      local.get 2
      call 49
      local.set 1
    end
    local.get 0
    local.get 1
    i32.store offset=4
    local.get 0
    local.get 3
    i32.store)
  (func (;49;) (type 4) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 91
    call 92)
  (func (;50;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 40
    unreachable)
  (func (;51;) (type 15) (param i64 i32 i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    local.get 4
    call 39
    call 6
    drop)
  (func (;52;) (type 16) (param i32 i32 i32 i32)
    local.get 0
    local.get 1
    i64.const 0
    local.get 2
    local.get 3
    call 7
    drop)
  (func (;53;) (type 3) (result i32)
    (local i32)
    call 39
    local.tee 0
    call 8
    local.get 0)
  (func (;54;) (type 13)
    (local i32)
    call 39
    local.tee 0
    call 9
    block  ;; label = @1
      local.get 0
      call 53
      call 47
      i32.eqz
      br_if 0 (;@1;)
      return
    end
    i32.const 131985
    i32.const 36
    call 2
    unreachable)
  (func (;55;) (type 3) (result i32)
    (local i32)
    block  ;; label = @1
      i32.const 0
      i32.load8_u offset=132132
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      i32.const -21
      i32.const 2147483647
      local.get 0
      select
      return
    end
    i32.const 0
    i32.const 1
    i32.store8 offset=132132
    i32.const -21
    call 10
    i32.const -21)
  (func (;56;) (type 7) (param i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i64 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 1
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        call 55
        local.tee 2
        call 11
        i32.const -16
        i32.and
        i32.const 16
        i32.ne
        br_if 0 (;@2;)
        i32.const 1
        local.set 3
        i32.const 0
        local.set 4
        local.get 2
        call 11
        local.set 5
        local.get 1
        i32.const 16
        i32.add
        local.set 6
        i32.const 0
        local.set 7
        loop  ;; label = @3
          local.get 7
          local.set 8
          local.get 4
          i32.const 16
          i32.add
          local.tee 9
          local.get 5
          i32.gt_u
          br_if 2 (;@1;)
          local.get 6
          i64.const 0
          i64.store
          local.get 1
          i64.const 0
          i64.store offset=8
          local.get 2
          local.get 4
          i32.const 16
          local.get 1
          i32.const 8
          i32.add
          call 12
          drop
          local.get 1
          i32.const 0
          i32.store offset=28
          i32.const 1
          local.set 7
          local.get 3
          i32.const 1
          i32.and
          local.set 10
          local.get 1
          i32.const 8
          i32.add
          local.get 1
          i32.const 28
          i32.add
          call 57
          local.set 11
          local.get 1
          i32.const 8
          i32.add
          local.get 1
          i32.const 28
          i32.add
          call 58
          local.set 12
          local.get 1
          i32.const 8
          i32.add
          local.get 1
          i32.const 28
          i32.add
          call 57
          local.set 13
          local.get 9
          local.set 4
          i32.const 0
          local.set 3
          local.get 10
          br_if 0 (;@3;)
        end
        local.get 8
        call 59
        unreachable
      end
      i32.const 131072
      i32.const 34
      call 2
      unreachable
    end
    block  ;; label = @1
      local.get 12
      i64.const 0
      i64.ne
      br_if 0 (;@1;)
      local.get 0
      local.get 13
      i32.store offset=4
      local.get 0
      local.get 11
      i32.store
      local.get 1
      i32.const 32
      i32.add
      global.set 0
      return
    end
    i32.const 131957
    i32.const 28
    call 2
    unreachable)
  (func (;57;) (type 4) (param i32 i32) (result i32)
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
      call 81
      unreachable
    end
    local.get 3
    call 82
    unreachable)
  (func (;58;) (type 17) (param i32 i32) (result i64)
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
      call 81
      unreachable
    end
    local.get 3
    call 82
    unreachable)
  (func (;59;) (type 7) (param i32)
    call 157
    unreachable)
  (func (;60;) (type 8) (param i32) (result i32)
    (local i32)
    local.get 0
    call 39
    local.tee 1
    call 13
    drop
    local.get 1)
  (func (;61;) (type 16) (param i32 i32 i32 i32)
    (local i32)
    i32.const 131106
    i32.const 23
    call 62
    local.tee 4
    local.get 0
    local.get 1
    call 14
    drop
    local.get 4
    i32.const 131129
    i32.const 3
    call 14
    drop
    local.get 4
    local.get 2
    local.get 3
    call 14
    drop
    local.get 4
    call 15
    unreachable)
  (func (;62;) (type 4) (param i32 i32) (result i32)
    (local i32)
    call 39
    local.tee 2
    local.get 0
    local.get 1
    call 23
    drop
    local.get 2)
  (func (;63;) (type 8) (param i32) (result i32)
    local.get 0
    call 60)
  (func (;64;) (type 3) (result i32)
    (local i32)
    block  ;; label = @1
      i32.const 1
      call 60
      local.tee 0
      call 11
      i32.const 32
      i32.eq
      br_if 0 (;@1;)
      i32.const 131425
      i32.const 12
      i32.const 131331
      i32.const 16
      call 61
      unreachable
    end
    local.get 0)
  (func (;65;) (type 3) (result i32)
    (local i32 i32 i32)
    i32.const 2
    call 60
    local.tee 0
    call 11
    local.set 1
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.const 0
        i32.const 32
        call 3
        local.tee 2
        call 16
        br_if 0 (;@2;)
        local.get 1
        i32.const 32
        i32.ne
        br_if 1 (;@1;)
        local.get 2
        return
      end
      i32.const 131423
      i32.const 2
      i32.const 131257
      i32.const 15
      call 61
      unreachable
    end
    i32.const 131423
    i32.const 2
    i32.const 131243
    i32.const 14
    call 61
    unreachable)
  (func (;66;) (type 4) (param i32 i32) (result i32)
    (local i64)
    block  ;; label = @1
      i32.const 1
      call 17
      local.tee 2
      i64.const 4294967296
      i64.lt_u
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      i32.const 131243
      i32.const 14
      call 61
      unreachable
    end
    local.get 2
    i32.wrap_i64)
  (func (;67;) (type 7) (param i32)
    block  ;; label = @1
      call 18
      local.get 0
      i32.ne
      br_if 0 (;@1;)
      return
    end
    i32.const 131132
    i32.const 25
    call 2
    unreachable)
  (func (;68;) (type 7) (param i32)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    local.get 0
    i32.load offset=4
    call 69
    i32.store offset=28
    local.get 1
    i32.const 1
    i32.store offset=24
    local.get 1
    local.get 0
    i32.store offset=20
    block  ;; label = @1
      loop  ;; label = @2
        local.get 1
        i32.const 8
        i32.add
        local.get 1
        i32.const 20
        i32.add
        call 48
        local.get 1
        i32.load offset=8
        i32.eqz
        br_if 1 (;@1;)
        local.get 1
        i32.load offset=12
        call 19
        drop
        br 0 (;@2;)
      end
    end
    local.get 1
    i32.const 32
    i32.add
    global.set 0)
  (func (;69;) (type 8) (param i32) (result i32)
    local.get 0
    call 93)
  (func (;70;) (type 16) (param i32 i32 i32 i32)
    local.get 2
    local.get 3
    call 71
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;71;) (type 2) (param i32 i32)
    call 43
    drop
    local.get 0
    local.get 1
    call 83
    call 80)
  (func (;72;) (type 16) (param i32 i32 i32 i32)
    local.get 2
    local.get 3
    call 73
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;73;) (type 2) (param i32 i32)
    call 43
    drop
    local.get 0
    local.get 1
    call 41
    call 80)
  (func (;74;) (type 1) (param i32 i32 i32)
    local.get 1
    local.get 2
    call 62
    local.set 2
    local.get 0
    call 43
    i32.store offset=4
    local.get 0
    local.get 2
    i32.store)
  (func (;75;) (type 2) (param i32 i32)
    (local i64)
    call 20
    local.set 2
    i32.const -30
    call 21
    local.get 2
    i32.const -30
    call 76
    local.get 0
    local.get 1
    call 51
    call 22)
  (func (;76;) (type 3) (result i32)
    i32.const -10
    i64.const 0
    call 0
    i32.const -10)
  (func (;77;) (type 1) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 14
    drop)
  (func (;78;) (type 16) (param i32 i32 i32 i32)
    local.get 0
    local.get 1
    local.get 3
    local.get 2
    call 12
    drop)
  (func (;79;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 4
    drop)
  (func (;80;) (type 2) (param i32 i32)
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
    call 14
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;81;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 159
    unreachable)
  (func (;82;) (type 7) (param i32)
    local.get 0
    call 158
    unreachable)
  (func (;83;) (type 8) (param i32) (result i32)
    (local i32)
    call 39
    local.tee 1
    local.get 0
    call 24
    drop
    local.get 1)
  (func (;84;) (type 18) (param i64) (result i32)
    block  ;; label = @1
      local.get 0
      i64.const -1
      i64.gt_s
      br_if 0 (;@1;)
      i32.const 131272
      i32.const 17
      call 2
      unreachable
    end
    i32.const -14
    local.get 0
    call 0
    i32.const -14)
  (func (;85;) (type 8) (param i32) (result i32)
    (local i32)
    local.get 0
    call 39
    local.tee 1
    call 25
    drop
    local.get 1)
  (func (;86;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 87)
  (func (;87;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 26
    drop)
  (func (;88;) (type 7) (param i32)
    i32.const -20
    i32.const 1
    i32.const 0
    call 23
    drop
    local.get 0
    i32.const -20
    call 26
    drop)
  (func (;89;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    i64.extend_i32_u
    call 90)
  (func (;90;) (type 0) (param i32 i64)
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
    call 117
    local.get 0
    local.get 2
    i32.load
    local.get 2
    i32.load offset=4
    call 112
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;91;) (type 4) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    call 41
    local.tee 0
    i32.const 131289
    i32.const 5
    call 14
    drop
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
    call 14
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func (;92;) (type 8) (param i32) (result i32)
    local.get 0
    call 85)
  (func (;93;) (type 8) (param i32) (result i32)
    (local i64)
    block  ;; label = @1
      local.get 0
      local.get 0
      call 98
      local.tee 1
      i64.const 4294967296
      i64.lt_u
      br_if 0 (;@1;)
      local.get 0
      i32.const 131243
      i32.const 14
      call 99
      unreachable
    end
    local.get 1
    i32.wrap_i64)
  (func (;94;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 83
    call 26
    drop)
  (func (;95;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    i64.extend_i32_u
    call 90)
  (func (;96;) (type 8) (param i32) (result i32)
    local.get 0
    call 85
    call 39
    local.tee 0
    call 27
    drop
    local.get 0)
  (func (;97;) (type 8) (param i32) (result i32)
    (local i64)
    block  ;; label = @1
      local.get 0
      local.get 0
      call 98
      local.tee 1
      i64.const 4294967296
      i64.lt_u
      br_if 0 (;@1;)
      local.get 0
      i32.const 131243
      i32.const 14
      call 99
      unreachable
    end
    local.get 1
    i32.wrap_i64)
  (func (;98;) (type 17) (param i32 i32) (result i64)
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
      call 85
      local.tee 3
      call 11
      local.tee 0
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 1
      i32.const 131243
      i32.const 14
      call 99
      unreachable
    end
    local.get 3
    i32.const 0
    local.get 0
    local.get 2
    i32.const 8
    i32.add
    local.get 0
    i32.sub
    i32.const 8
    i32.add
    call 12
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
  (func (;99;) (type 1) (param i32 i32 i32)
    (local i32)
    i32.const 131300
    i32.const 27
    call 62
    local.tee 3
    local.get 0
    call 4
    drop
    local.get 3
    i32.const 131129
    i32.const 3
    call 14
    drop
    local.get 3
    local.get 1
    local.get 2
    call 14
    drop
    local.get 3
    call 15
    unreachable)
  (func (;100;) (type 8) (param i32) (result i32)
    (local i64)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          local.get 0
          call 98
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
        i32.const 131347
        i32.const 18
        call 99
        unreachable
      end
      i32.const 1
      local.set 0
    end
    local.get 0)
  (func (;101;) (type 8) (param i32) (result i32)
    local.get 0
    i32.const -25
    call 25
    drop
    i32.const -25
    call 11
    i32.eqz)
  (func (;102;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 103
    call 88)
  (func (;103;) (type 4) (param i32 i32) (result i32)
    local.get 0
    call 41
    local.tee 0
    i32.const 131294
    i32.const 6
    call 14
    drop
    local.get 1
    local.get 0
    call 109
    local.get 0)
  (func (;104;) (type 2) (param i32 i32)
    (local i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      local.get 0
      i32.load offset=8
      local.tee 2
      local.get 1
      call 105
      local.tee 3
      i32.eqz
      br_if 0 (;@1;)
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              i32.load offset=4
              local.tee 4
              call 69
              local.tee 5
              local.get 3
              i32.lt_u
              br_if 0 (;@5;)
              local.get 5
              local.get 3
              i32.eq
              br_if 3 (;@2;)
              local.get 4
              call 69
              local.get 5
              i32.lt_u
              br_if 1 (;@4;)
              local.get 0
              i32.load
              local.tee 6
              local.get 5
              call 49
              local.set 7
              local.get 4
              call 69
              local.get 3
              i32.lt_u
              br_if 2 (;@3;)
              local.get 6
              local.get 3
              call 91
              local.get 7
              call 86
              br 3 (;@2;)
            end
            i32.const 132021
            i32.const 18
            call 2
            unreachable
          end
          i32.const 132021
          i32.const 18
          call 2
          unreachable
        end
        i32.const 132021
        i32.const 18
        call 2
        unreachable
      end
      block  ;; label = @2
        local.get 4
        call 69
        local.get 5
        i32.ge_u
        br_if 0 (;@2;)
        i32.const 132021
        i32.const 18
        call 2
        unreachable
      end
      local.get 0
      i32.load
      local.get 5
      call 91
      call 88
      local.get 4
      local.get 5
      i32.const -1
      i32.add
      call 89
      block  ;; label = @2
        local.get 5
        local.get 3
        i32.eq
        br_if 0 (;@2;)
        local.get 2
        local.get 7
        local.get 3
        call 106
      end
      local.get 2
      local.get 1
      call 102
    end)
  (func (;105;) (type 4) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 103
    call 93)
  (func (;106;) (type 1) (param i32 i32 i32)
    local.get 0
    local.get 1
    call 103
    local.get 2
    i64.extend_i32_u
    call 90)
  (func (;107;) (type 2) (param i32 i32)
    (local i32 i32 i32)
    block  ;; label = @1
      local.get 0
      i32.load offset=8
      local.tee 2
      local.get 1
      call 108
      br_if 0 (;@1;)
      local.get 0
      i32.load offset=4
      local.tee 3
      call 69
      local.set 4
      local.get 0
      i32.load
      local.get 4
      i32.const 1
      i32.add
      local.tee 0
      call 91
      local.get 1
      call 86
      local.get 3
      local.get 0
      call 89
      local.get 2
      local.get 1
      local.get 3
      call 69
      call 106
    end)
  (func (;108;) (type 4) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call 105
    i32.const 0
    i32.ne)
  (func (;109;) (type 2) (param i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    call 44
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
    call 77
    local.get 1
    local.get 0
    call 79
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;110;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 71)
  (func (;111;) (type 4) (param i32 i32) (result i32)
    (local i32)
    call 43
    local.tee 2
    local.get 0
    local.get 1
    call 62
    call 80
    local.get 2)
  (func (;112;) (type 1) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 62
    call 26
    drop)
  (func (;113;) (type 2) (param i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    call 44
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
    call 14
    drop
    local.get 1
    local.get 0
    call 79
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;114;) (type 2) (param i32 i32)
    call 43
    drop
    local.get 0
    local.get 1
    call 41
    call 80)
  (func (;115;) (type 2) (param i32 i32)
    (local i32 i32)
    local.get 1
    call 41
    local.set 2
    local.get 1
    call 41
    local.tee 3
    i32.const 131327
    i32.const 4
    call 14
    drop
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store
    local.get 0
    local.get 2
    i32.store offset=8)
  (func (;116;) (type 19) (param i64 i32)
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
    call 77
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;117;) (type 20) (param i32 i64 i32 i32)
    (local i32 i32 i32)
    local.get 3
    local.get 1
    i64.store8 offset=7
    local.get 3
    local.get 1
    i64.const 24
    i64.shr_u
    i64.store8 offset=4
    local.get 3
    i32.const 0
    i32.store align=1
    local.get 3
    local.get 1
    i64.const 8
    i64.shr_u
    i32.wrap_i64
    local.tee 4
    i32.store8 offset=6
    local.get 3
    local.get 1
    i64.const 16
    i64.shr_u
    i32.wrap_i64
    local.tee 5
    i32.store8 offset=5
    local.get 0
    i32.const 8
    i32.const 5
    i32.const 4
    local.get 1
    i64.const 16777216
    i64.lt_u
    local.tee 6
    select
    local.get 6
    local.get 5
    i32.const 255
    i32.and
    i32.eqz
    i32.and
    local.tee 5
    i32.add
    i32.const 0
    local.get 5
    local.get 4
    i32.const 255
    i32.and
    select
    local.tee 4
    i32.add
    local.get 4
    i32.const 0
    local.get 1
    i64.eqz
    select
    i32.add
    local.tee 4
    local.get 3
    local.get 4
    i32.const 7
    i32.and
    i32.add
    i32.load8_s
    i32.const 0
    i32.lt_s
    local.get 2
    i32.and
    i32.sub
    local.tee 2
    i32.sub
    i32.store offset=4
    local.get 0
    local.get 3
    local.get 2
    i32.add
    i32.store)
  (func (;118;) (type 3) (result i32)
    call 119
    call 100
    i32.const 1
    i32.xor)
  (func (;119;) (type 3) (result i32)
    i32.const 131791
    i32.const 19
    call 62)
  (func (;120;) (type 7) (param i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    call 119
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
    call 117
    local.get 2
    local.get 1
    i32.load
    local.get 1
    i32.load offset=4
    call 112
    local.get 1
    i32.const 16
    i32.add
    global.set 0)
  (func (;121;) (type 8) (param i32) (result i32)
    (local i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 1
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              call 118
              i32.eqz
              br_if 0 (;@5;)
              local.get 1
              i32.const 32
              i32.add
              call 56
              local.get 1
              i32.load offset=36
              local.set 2
              local.get 1
              local.get 1
              i32.load offset=32
              local.tee 3
              i32.store offset=44
              local.get 2
              call 122
              i32.const 255
              i32.and
              i32.const 1
              i32.le_u
              br_if 1 (;@4;)
              local.get 3
              local.get 0
              call 123
              call 92
              local.tee 4
              call 47
              i32.eqz
              br_if 2 (;@3;)
              local.get 1
              i32.const 44
              i32.add
              local.get 0
              call 124
              local.get 1
              i32.const 44
              i32.add
              local.get 0
              local.get 2
              call 38
              call 125
              local.tee 5
              local.get 0
              call 126
              local.tee 6
              call 96
              local.tee 3
              call 28
              i32.const 1
              i32.ge_s
              br_if 3 (;@2;)
              local.get 3
              local.get 3
              local.get 5
              call 29
              local.get 3
              call 122
              i32.const 255
              i32.and
              i32.eqz
              br_if 4 (;@1;)
              local.get 6
              local.get 3
              call 94
              local.get 1
              i32.const 24
              i32.add
              i32.const 131205
              i32.const 13
              call 74
              local.get 1
              i32.const 16
              i32.add
              local.get 1
              i32.load offset=24
              local.get 1
              i32.load offset=28
              local.get 4
              call 72
              local.get 1
              i32.const 8
              i32.add
              local.get 1
              i32.load offset=16
              local.get 1
              i32.load offset=20
              local.get 2
              call 70
              local.get 1
              i32.load offset=8
              local.get 1
              i32.load offset=12
              call 75
              local.get 5
              call 38
              local.set 2
              i32.const 131944
              i32.const 13
              call 111
              local.tee 3
              local.get 0
              call 127
              local.get 3
              local.get 2
              call 110
              local.get 3
              call 43
              call 30
              local.get 1
              i32.const 48
              i32.add
              global.set 0
              local.get 5
              return
            end
            i32.const 131437
            i32.const 18
            call 50
            unreachable
          end
          i32.const 131520
          i32.const 28
          call 50
          unreachable
        end
        i32.const 131548
        i32.const 22
        call 50
        unreachable
      end
      i32.const 131570
      i32.const 35
      call 50
      unreachable
    end
    i32.const 131157
    i32.const 48
    call 2
    unreachable)
  (func (;122;) (type 8) (param i32) (result i32)
    i32.const 0
    i32.const 2
    i32.const 1
    local.get 0
    call 37
    local.tee 0
    select
    local.get 0
    i32.const 0
    i32.lt_s
    select)
  (func (;123;) (type 8) (param i32) (result i32)
    (local i32)
    i32.const 131902
    i32.const 31
    call 62
    local.set 1
    local.get 0
    i32.load
    local.get 1
    call 109
    local.get 1)
  (func (;124;) (type 2) (param i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        call 128
        call 101
        br_if 0 (;@2;)
        local.get 1
        call 128
        call 101
        i32.eqz
        br_if 1 (;@1;)
        i32.const 131707
        i32.const 38
        call 50
        unreachable
      end
      i32.const 131745
      i32.const 33
      call 50
      unreachable
    end)
  (func (;125;) (type 10) (param i32 i32 i32) (result i32)
    block  ;; label = @1
      local.get 0
      call 128
      call 97
      local.tee 0
      local.get 1
      call 128
      call 97
      local.tee 1
      i32.eq
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 0
        local.get 1
        i32.gt_u
        br_if 0 (;@2;)
        local.get 2
        local.get 2
        local.get 1
        local.get 0
        i32.sub
        call 129
        call 84
        call 31
        br 1 (;@1;)
      end
      local.get 2
      local.get 2
      local.get 0
      local.get 1
      i32.sub
      call 129
      call 84
      call 32
    end
    local.get 2
    call 38)
  (func (;126;) (type 8) (param i32) (result i32)
    (local i32)
    i32.const 131825
    i32.const 14
    call 62
    local.set 1
    local.get 0
    i32.load
    local.get 1
    call 109
    local.get 1)
  (func (;127;) (type 2) (param i32 i32)
    call 43
    drop
    local.get 0
    local.get 1
    i32.load
    call 41
    call 80)
  (func (;128;) (type 8) (param i32) (result i32)
    (local i32)
    i32.const 131839
    i32.const 18
    call 62
    local.set 1
    local.get 0
    i32.load
    local.get 1
    call 109
    local.get 1)
  (func (;129;) (type 11) (param i32) (result i64)
    (local i64 i64)
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        br_if 0 (;@2;)
        i64.const 1
        local.set 1
        br 1 (;@1;)
      end
      i64.const 10
      local.set 2
      i64.const 1
      local.set 1
      loop  ;; label = @2
        block  ;; label = @3
          local.get 0
          i32.const 1
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 2
          local.get 1
          i64.mul
          local.set 1
          local.get 0
          i32.const 1
          i32.eq
          br_if 2 (;@1;)
        end
        local.get 0
        i32.const 1
        i32.shr_u
        local.set 0
        local.get 2
        local.get 2
        i64.mul
        local.set 2
        br 0 (;@2;)
      end
    end
    local.get 1)
  (func (;130;) (type 7) (param i32)
    (local i64)
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        call 33
        local.tee 1
        i32.const 1
        call 131
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        i32.const 2
        call 131
        br_if 1 (;@1;)
      end
      i32.const 131682
      i32.const 25
      call 50
      unreachable
    end)
  (func (;131;) (type 21) (param i64 i32) (result i32)
    local.get 1
    i32.extend8_s
    i32.const 3
    i32.shl
    i32.const 132056
    i32.add
    i64.load
    local.get 0
    i64.and
    i64.const 0
    i64.ne)
  (func (;132;) (type 7) (param i32)
    (local i32)
    call 43
    local.set 1
    local.get 0
    call 43
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;133;) (type 2) (param i32 i32)
    (local i32)
    i32.const 131857
    i32.const 21
    call 62
    local.set 2
    local.get 1
    i32.load
    local.get 2
    call 109
    local.get 0
    local.get 2
    call 115)
  (func (;134;) (type 7) (param i32)
    local.get 0
    i32.const 131878
    i32.const 24
    call 62
    call 115)
  (func (;135;) (type 13)
    call 136
    unreachable)
  (func (;136;) (type 13)
    i32.const 132039
    i32.const 14
    call 2
    unreachable)
  (func (;137;) (type 13)
    call 34
    i32.const 0
    call 67
    i32.const 1
    call 120)
  (func (;138;) (type 13)
    call 34
    i32.const 0
    call 67)
  (func (;139;) (type 13)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 34
    call 54
    i32.const 2
    call 67
    i32.const 0
    call 63
    local.set 1
    i32.const 131411
    i32.const 12
    call 66
    local.set 2
    local.get 0
    local.get 1
    i32.store
    local.get 1
    call 130
    local.get 0
    call 128
    local.get 2
    call 95
    local.get 0
    i32.const 4
    i32.add
    call 134
    local.get 0
    i32.const 4
    i32.add
    local.get 1
    call 107
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;140;) (type 13)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 34
    call 54
    i32.const 2
    call 67
    i32.const 0
    call 63
    local.set 1
    i32.const 131411
    i32.const 12
    call 66
    local.set 2
    local.get 0
    local.get 1
    i32.store
    local.get 0
    i32.const 4
    i32.add
    call 134
    block  ;; label = @1
      local.get 0
      i32.load offset=12
      local.get 1
      call 108
      br_if 0 (;@1;)
      i32.const 131605
      i32.const 33
      call 50
      unreachable
    end
    local.get 0
    call 128
    local.get 2
    call 95
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;141;) (type 13)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 0
    global.set 0
    call 34
    call 54
    i32.const 1
    call 67
    local.get 0
    i32.const 0
    call 63
    local.tee 1
    i32.store offset=16
    local.get 0
    i32.const 36
    i32.add
    call 134
    local.get 0
    i32.const 36
    i32.add
    local.get 1
    call 104
    local.get 0
    i32.const 20
    i32.add
    local.get 0
    i32.const 16
    i32.add
    call 133
    local.get 0
    local.get 0
    i32.load offset=24
    call 93
    i32.store offset=44
    local.get 0
    i32.const 1
    i32.store offset=40
    local.get 0
    local.get 0
    i32.const 20
    i32.add
    i32.store offset=36
    block  ;; label = @1
      loop  ;; label = @2
        local.get 0
        i32.const 8
        i32.add
        local.get 0
        i32.const 36
        i32.add
        call 48
        block  ;; label = @3
          local.get 0
          i32.load offset=8
          br_if 0 (;@3;)
          local.get 0
          local.get 0
          i32.load offset=24
          call 93
          i32.store offset=44
          local.get 0
          i32.const 1
          i32.store offset=40
          local.get 0
          local.get 0
          i32.const 20
          i32.add
          i32.store offset=36
          loop  ;; label = @4
            local.get 0
            local.get 0
            i32.const 36
            i32.add
            call 48
            block  ;; label = @5
              local.get 0
              i32.load
              br_if 0 (;@5;)
              i32.const 0
              local.set 2
              local.get 0
              i32.load offset=24
              call 93
              local.set 3
              i32.const 1
              local.set 1
              loop  ;; label = @6
                local.get 2
                i32.const 1
                i32.and
                br_if 5 (;@1;)
                local.get 1
                local.get 3
                i32.gt_u
                br_if 5 (;@1;)
                local.get 0
                i32.load offset=20
                local.get 1
                call 91
                call 88
                local.get 1
                local.get 3
                i32.ge_u
                local.set 2
                local.get 1
                local.get 1
                local.get 3
                i32.lt_u
                i32.add
                local.set 1
                br 0 (;@6;)
              end
            end
            local.get 0
            i32.load offset=28
            local.get 0
            i32.load offset=4
            call 102
            br 0 (;@4;)
          end
        end
        local.get 0
        local.get 0
        i32.load offset=12
        i32.store offset=32
        local.get 0
        i32.const 32
        i32.add
        call 123
        call 88
        local.get 0
        i32.const 32
        i32.add
        call 128
        call 88
        br 0 (;@2;)
      end
    end
    local.get 0
    i32.load offset=24
    i64.const 0
    call 90
    local.get 0
    i32.const 16
    i32.add
    call 128
    call 88
    local.get 0
    i32.const 48
    i32.add
    global.set 0)
  (func (;142;) (type 13)
    (local i32 i32 i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    call 34
    call 54
    i32.const 3
    call 67
    i32.const 0
    call 63
    local.set 1
    i32.const 131382
    i32.const 29
    call 66
    local.set 2
    local.get 0
    i32.const 2
    call 63
    local.tee 3
    i32.store offset=16
    local.get 0
    local.get 1
    i32.store offset=12
    local.get 3
    call 130
    block  ;; label = @1
      local.get 0
      i32.const 12
      i32.add
      call 123
      local.tee 4
      call 101
      br_if 0 (;@1;)
      i32.const 131455
      i32.const 65
      call 50
      unreachable
    end
    local.get 0
    i32.const 12
    i32.add
    call 128
    local.get 2
    call 95
    local.get 4
    local.get 3
    call 86
    local.get 0
    i32.const 20
    i32.add
    local.get 0
    i32.const 16
    i32.add
    call 133
    local.get 0
    i32.const 20
    i32.add
    local.get 1
    call 107
    local.get 0
    i32.const 20
    i32.add
    call 134
    local.get 0
    i32.const 20
    i32.add
    local.get 3
    call 107
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;143;) (type 13)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 34
    call 54
    i32.const 2
    call 67
    i32.const 0
    call 63
    local.set 1
    i32.const 131382
    i32.const 29
    call 66
    local.set 2
    local.get 0
    local.get 1
    i32.store offset=12
    block  ;; label = @1
      local.get 0
      i32.const 12
      i32.add
      call 123
      call 101
      i32.eqz
      br_if 0 (;@1;)
      i32.const 131638
      i32.const 44
      call 50
      unreachable
    end
    local.get 0
    i32.const 12
    i32.add
    call 128
    local.get 2
    call 95
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;144;) (type 13)
    (local i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    call 34
    call 54
    i32.const 1
    call 67
    local.get 0
    i32.const 0
    call 63
    local.tee 1
    i32.store offset=12
    local.get 0
    local.get 0
    i32.const 12
    i32.add
    call 123
    local.tee 2
    call 92
    i32.store offset=16
    local.get 0
    i32.const 20
    i32.add
    local.get 0
    i32.const 16
    i32.add
    call 133
    local.get 0
    i32.const 20
    i32.add
    local.get 1
    call 104
    local.get 2
    call 88
    local.get 0
    i32.const 12
    i32.add
    call 128
    call 88
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;145;) (type 13)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    i32.const 0
    call 67
    local.get 0
    call 56
    local.get 0
    i32.load offset=4
    local.set 1
    local.get 0
    local.get 0
    i32.load
    i32.store offset=12
    local.get 0
    i32.const 12
    i32.add
    call 126
    local.tee 2
    call 96
    local.tee 3
    local.get 3
    local.get 1
    call 1
    local.get 2
    local.get 3
    call 94
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;146;) (type 13)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i64)
    global.get 0
    i32.const 128
    i32.sub
    local.tee 0
    global.set 0
    i32.const 0
    call 67
    block  ;; label = @1
      block  ;; label = @2
        call 118
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        call 55
        i32.store offset=104
        local.get 0
        local.get 0
        i32.const 104
        i32.add
        call 42
        local.tee 1
        i32.store offset=44
        local.get 1
        call 11
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        call 43
        local.tee 2
        i32.store offset=48
        local.get 0
        local.get 1
        call 11
        i32.store offset=60
        local.get 0
        i32.const 0
        i32.store offset=56
        local.get 0
        local.get 0
        i32.const 44
        i32.add
        i32.store offset=52
        local.get 0
        i32.const 64
        i32.add
        i32.const 8
        i32.add
        local.set 3
        local.get 0
        i32.const 104
        i32.add
        i32.const 8
        i32.add
        local.set 1
        loop  ;; label = @3
          local.get 0
          i32.const 104
          i32.add
          local.get 0
          i32.const 52
          i32.add
          call 45
          block  ;; label = @4
            local.get 0
            i32.load offset=104
            br_if 0 (;@4;)
            local.get 0
            i32.const 48
            i32.add
            call 42
            local.set 3
            i32.const -30
            call 8
            local.get 0
            i32.const 32
            i32.add
            call 132
            i32.const -30
            local.get 3
            local.get 0
            i32.load offset=32
            local.get 0
            i32.load offset=36
            call 52
            local.get 0
            i32.load offset=48
            local.set 1
            br 3 (;@1;)
          end
          local.get 3
          local.get 1
          i32.const 8
          i32.add
          i64.load
          i64.store
          local.get 0
          local.get 1
          i64.load
          i64.store offset=64
          block  ;; label = @4
            local.get 3
            call 123
            local.tee 4
            call 101
            br_if 0 (;@4;)
            local.get 0
            local.get 4
            call 92
            local.tee 4
            i32.store offset=84
            local.get 0
            i32.const 84
            i32.add
            local.get 3
            call 124
            local.get 3
            call 126
            local.tee 5
            call 96
            local.tee 6
            local.get 6
            local.get 0
            i32.load offset=76
            local.tee 7
            call 1
            local.get 5
            local.get 6
            call 94
            local.get 3
            local.get 0
            i32.const 84
            i32.add
            local.get 7
            call 125
            local.set 6
            local.get 0
            i32.const 24
            i32.add
            i32.const 131218
            i32.const 13
            call 74
            local.get 0
            i32.const 16
            i32.add
            local.get 0
            i32.load offset=24
            local.get 0
            i32.load offset=28
            local.get 4
            call 72
            local.get 0
            i32.const 8
            i32.add
            local.get 0
            i32.load offset=16
            local.get 0
            i32.load offset=20
            local.get 6
            call 70
            local.get 0
            i32.load offset=8
            local.get 0
            i32.load offset=12
            call 75
            local.get 4
            call 41
            local.set 5
            local.get 0
            local.get 6
            call 38
            i32.store offset=100
            local.get 0
            i64.const 0
            i64.store offset=88
            local.get 0
            local.get 5
            i32.store offset=96
            local.get 2
            local.get 0
            i32.const 88
            i32.add
            call 46
            i32.const 131933
            i32.const 11
            call 111
            local.tee 5
            local.get 4
            call 73
            local.get 5
            local.get 6
            call 110
            local.get 5
            call 43
            call 30
            br 1 (;@3;)
          end
          local.get 2
          local.get 1
          call 46
          br 0 (;@3;)
        end
      end
      i32.const 131437
      i32.const 18
      call 50
      unreachable
    end
    local.get 0
    local.get 1
    i32.store offset=52
    i32.const 1
    i32.const 0
    call 62
    local.set 3
    local.get 0
    local.get 1
    call 11
    i32.store offset=96
    local.get 0
    i32.const 0
    i32.store offset=92
    local.get 0
    local.get 0
    i32.const 52
    i32.add
    i32.store offset=88
    block  ;; label = @1
      loop  ;; label = @2
        local.get 0
        i32.const 104
        i32.add
        local.get 0
        i32.const 88
        i32.add
        call 45
        local.get 0
        i32.load offset=104
        i32.const 1
        i32.ne
        br_if 1 (;@1;)
        local.get 0
        i32.load offset=124
        local.set 1
        local.get 0
        i64.load offset=112
        local.set 8
        local.get 0
        i32.load offset=120
        local.get 3
        call 113
        local.get 0
        local.get 8
        i64.const 56
        i64.shl
        local.get 8
        i64.const 65280
        i64.and
        i64.const 40
        i64.shl
        i64.or
        local.get 8
        i64.const 16711680
        i64.and
        i64.const 24
        i64.shl
        local.get 8
        i64.const 4278190080
        i64.and
        i64.const 8
        i64.shl
        i64.or
        i64.or
        local.get 8
        i64.const 8
        i64.shr_u
        i64.const 4278190080
        i64.and
        local.get 8
        i64.const 24
        i64.shr_u
        i64.const 16711680
        i64.and
        i64.or
        local.get 8
        i64.const 40
        i64.shr_u
        i64.const 65280
        i64.and
        local.get 8
        i64.const 56
        i64.shr_u
        i64.or
        i64.or
        i64.or
        i64.store offset=64
        local.get 3
        local.get 0
        i32.const 64
        i32.add
        i32.const 8
        call 14
        drop
        local.get 1
        call 83
        local.get 3
        call 113
        br 0 (;@2;)
      end
    end
    local.get 3
    call 19
    drop
    local.get 0
    i32.const 128
    i32.add
    global.set 0)
  (func (;147;) (type 13)
    (local i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    i32.const 1
    call 67
    local.get 0
    i32.const 0
    call 63
    i32.store offset=12
    local.get 0
    i32.const 12
    i32.add
    call 121
    local.set 1
    i32.const -30
    call 8
    local.get 0
    call 132
    local.get 0
    i32.load offset=4
    local.set 2
    local.get 0
    i32.load
    local.set 3
    call 43
    local.set 4
    local.get 0
    i32.load offset=12
    call 41
    local.set 5
    local.get 0
    local.get 1
    call 38
    i32.store offset=28
    local.get 0
    i64.const 0
    i64.store offset=16
    local.get 0
    local.get 5
    i32.store offset=24
    local.get 4
    local.get 0
    i32.const 16
    i32.add
    call 46
    i32.const -30
    local.get 4
    local.get 3
    local.get 2
    call 52
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;148;) (type 13)
    (local i32 i32 i32 i32 i32 i32 i32 i64 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    i32.const 3
    call 67
    i32.const 0
    call 63
    local.set 1
    call 64
    local.set 2
    call 65
    local.set 3
    local.get 0
    local.get 1
    i32.store offset=20
    local.get 0
    local.get 0
    i32.const 20
    i32.add
    call 121
    i32.store offset=24
    call 53
    local.set 4
    local.get 0
    i32.const 8
    i32.add
    i32.const 131365
    i32.const 17
    call 74
    local.get 0
    i32.load offset=8
    local.set 5
    local.get 0
    i32.load offset=12
    local.set 6
    call 43
    call 41
    local.tee 1
    local.get 3
    call 4
    drop
    local.get 6
    local.get 1
    call 80
    call 43
    call 41
    local.tee 1
    local.get 4
    call 4
    drop
    i64.const 0
    local.get 1
    call 116
    i64.const 0
    local.get 1
    call 116
    local.get 6
    local.get 1
    call 80
    call 20
    local.set 7
    local.get 0
    i32.const 131231
    i32.const 12
    call 74
    local.get 0
    i32.load
    local.set 8
    local.get 0
    i32.load offset=4
    local.tee 3
    local.get 0
    i32.const 20
    i32.add
    call 127
    call 43
    drop
    local.get 3
    local.get 0
    i32.load offset=24
    call 83
    call 80
    block  ;; label = @1
      local.get 5
      call 11
      i32.eqz
      br_if 0 (;@1;)
      local.get 3
      local.get 5
      call 114
      local.get 6
      call 11
      local.set 4
      i32.const 0
      local.set 1
      loop  ;; label = @2
        local.get 1
        i32.const 4
        i32.add
        local.tee 5
        local.get 4
        i32.gt_u
        br_if 1 (;@1;)
        local.get 0
        i32.const 0
        i32.store offset=28
        local.get 6
        local.get 1
        i32.const 4
        local.get 0
        i32.const 28
        i32.add
        call 12
        drop
        local.get 3
        local.get 0
        i32.load offset=28
        local.tee 1
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
        call 114
        local.get 5
        local.set 1
        br 0 (;@2;)
      end
    end
    local.get 7
    i64.const -100000
    i64.add
    local.get 7
    local.get 7
    i64.const 100000
    i64.gt_u
    select
    local.get 2
    call 76
    local.get 8
    local.get 3
    call 51
    call 22
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;149;) (type 13)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 34
    i32.const 0
    call 67
    local.get 0
    i32.const 4
    i32.add
    call 134
    local.get 0
    i32.const 4
    i32.add
    call 68
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;150;) (type 13)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 34
    i32.const 1
    call 67
    local.get 0
    i32.const 0
    call 63
    i32.store offset=12
    local.get 0
    i32.const 12
    i32.add
    call 126
    call 96
    call 35
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;151;) (type 13)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 34
    i32.const 1
    call 67
    local.get 0
    i32.const 0
    call 63
    i32.store offset=12
    local.get 0
    i32.const 12
    i32.add
    call 123
    call 92
    call 19
    drop
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;152;) (type 13)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 34
    i32.const 1
    call 67
    local.get 0
    i32.const 0
    call 63
    i32.store
    local.get 0
    i32.const 4
    i32.add
    local.get 0
    call 133
    local.get 0
    i32.const 4
    i32.add
    call 68
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;153;) (type 13)
    call 34
    call 54
    i32.const 0
    call 67
    i32.const 1
    call 120
    i32.const 131778
    i32.const 13
    call 111
    call 43
    call 30)
  (func (;154;) (type 13)
    call 34
    call 54
    i32.const 0
    call 67
    i32.const 0
    call 120
    i32.const 131810
    i32.const 15
    call 111
    call 43
    call 30)
  (func (;155;) (type 13)
    call 34
    i32.const 0
    call 67
    call 119
    call 100
    i64.extend_i32_u
    call 36)
  (func (;156;) (type 13))
  (func (;157;) (type 13)
    call 135
    unreachable)
  (func (;158;) (type 7) (param i32)
    call 157
    unreachable)
  (func (;159;) (type 2) (param i32 i32)
    call 157
    unreachable)
  (memory (;0;) 3)
  (global (;0;) (mut i32) (i32.const 131072))
  (global (;1;) i32 (i32.const 132133))
  (global (;2;) i32 (i32.const 132144))
  (export "memory" (memory 0))
  (export "init" (func 137))
  (export "upgrade" (func 138))
  (export "addWrappedToken" (func 139))
  (export "updateWrappedToken" (func 140))
  (export "removeWrappedToken" (func 141))
  (export "whitelistToken" (func 142))
  (export "updateWhitelistedToken" (func 143))
  (export "blacklistToken" (func 144))
  (export "depositLiquidity" (func 145))
  (export "wrapTokens" (func 146))
  (export "unwrapToken" (func 147))
  (export "unwrapTokenCreateTransaction" (func 148))
  (export "getUniversalBridgedTokenIds" (func 149))
  (export "getTokenLiquidity" (func 150))
  (export "getChainSpecificToUniversalMapping" (func 151))
  (export "getchainSpecificTokenIds" (func 152))
  (export "pause" (func 153))
  (export "unpause" (func 154))
  (export "isPaused" (func 155))
  (export "callBack" (func 156))
  (export "__data_end" (global 1))
  (export "__heap_base" (global 2))
  (data (;0;) (i32.const 131072) "incorrect number of DCDT transfersargument decode error (): wrong number of argumentscannot subtract because result would be negativeDCDTLocalBurnDCDTLocalMintDCDTTransferinput too longinput too shortcast to i64 error.item.indexstorage decode error (key: .lenbad array lengthinput out of rangecreateTransactionchain_specific_token_decimalsnum_decimalstosafe_addressContract is pausedChain-specific token is already mapped to another universal tokenMust pay more than 0 tokens!Dcdt token unavailableContract does not have enough fundsUniversal token was not added yetChain-specific token was not whitelisted yetMust set local role firstChain-specific token requires updatingUniversal token requires updatingpauseContractpause_module:pausedunpauseContracttokenLiquiditytoken_decimals_numchainSpecificTokenIdsuniversalBridgedTokenIdschainSpecificToUniversalMappingwrap_tokensunwrap_tokensfungible DCDT token expectedEndpoint can only be called by ownerindex out of rangepanic occurred\00\00\00\00\00\00\00\00\00\00\00\01\00\00\00\00\00\00\00\02\00\00\00\00\00\00\00\04\00\00\00\00\00\00\00\08\00\00\00\00\00\00\00\10\00\00\00\00\00\00\00 \00\00\00\00\00\00\00@\00\00\00\00\00\00\00\80\00\00\00\00\00\00\00")
  (data (;1;) (i32.const 132128) "8\ff\ff\ff"))
