(module
  (type (;0;) (func (param i32 i64)))
  (type (;1;) (func (param i32 i32 i32)))
  (type (;2;) (func (param i32 i32)))
  (type (;3;) (func (result i32)))
  (type (;4;) (func (param i32 i32) (result i32)))
  (type (;5;) (func (param i64 i32 i32 i32 i32 i32) (result i32)))
  (type (;6;) (func (param i32)))
  (type (;7;) (func (param i32) (result i32)))
  (type (;8;) (func (param i32 i32 i32) (result i32)))
  (type (;9;) (func (param i32) (result i64)))
  (type (;10;) (func (param i64)))
  (type (;11;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;12;) (func (result i64)))
  (type (;13;) (func (param i64) (result i32)))
  (type (;14;) (func (param i32 i32 i64 i32 i32 i32 i32 i32 i32 i32 i32)))
  (type (;15;) (func))
  (type (;16;) (func (param i32 i32 i64 i32 i32) (result i32)))
  (type (;17;) (func (param i64 i32 i32 i32 i32) (result i32)))
  (type (;18;) (func (param i32 i32 i32 i32)))
  (type (;19;) (func (param i32 i32 i32 i32 i32)))
  (type (;20;) (func (param i32 i32) (result i64)))
  (type (;21;) (func (param i32 i32 i32 i64)))
  (type (;22;) (func (param i64 i32 i32 i32 i32)))
  (type (;23;) (func (param i64 i32)))
  (type (;24;) (func (param i32 i64 i64) (result i32)))
  (type (;25;) (func (param i32 i32 i32) (result i64)))
  (type (;26;) (func (param i32 i64 i32)))
  (import "env" "bigIntSetInt64" (func (;0;) (type 0)))
  (import "env" "bigIntAdd" (func (;1;) (type 1)))
  (import "env" "signalError" (func (;2;) (type 2)))
  (import "env" "mBufferNew" (func (;3;) (type 3)))
  (import "env" "mBufferAppend" (func (;4;) (type 4)))
  (import "env" "managedExecuteOnDestContext" (func (;5;) (type 5)))
  (import "env" "managedSCAddress" (func (;6;) (type 6)))
  (import "env" "isSmartContract" (func (;7;) (type 7)))
  (import "env" "getShardOfAddress" (func (;8;) (type 7)))
  (import "env" "managedOwnerAddress" (func (;9;) (type 6)))
  (import "env" "managedCaller" (func (;10;) (type 6)))
  (import "env" "mBufferEq" (func (;11;) (type 4)))
  (import "env" "mBufferGetArgument" (func (;12;) (type 4)))
  (import "env" "mBufferAppendBytes" (func (;13;) (type 8)))
  (import "env" "managedSignalError" (func (;14;) (type 6)))
  (import "env" "mBufferGetLength" (func (;15;) (type 7)))
  (import "env" "bigIntGetUnsignedArgument" (func (;16;) (type 2)))
  (import "env" "smallIntGetUnsignedArgument" (func (;17;) (type 9)))
  (import "env" "getNumArguments" (func (;18;) (type 3)))
  (import "env" "smallIntFinishUnsigned" (func (;19;) (type 10)))
  (import "env" "mBufferFinish" (func (;20;) (type 7)))
  (import "env" "mBufferSetBytes" (func (;21;) (type 8)))
  (import "env" "mBufferCopyByteSlice" (func (;22;) (type 11)))
  (import "env" "mBufferFromBigIntUnsigned" (func (;23;) (type 4)))
  (import "env" "mBufferToBigIntUnsigned" (func (;24;) (type 4)))
  (import "env" "mBufferStorageLoad" (func (;25;) (type 4)))
  (import "env" "mBufferStorageStore" (func (;26;) (type 4)))
  (import "env" "getBlockNonce" (func (;27;) (type 12)))
  (import "env" "bigIntNew" (func (;28;) (type 13)))
  (import "env" "managedGetDCDTTokenData" (func (;29;) (type 14)))
  (import "env" "getGasLeft" (func (;30;) (type 12)))
  (import "env" "checkNoPayment" (func (;31;) (type 15)))
  (import "env" "cleanReturnData" (func (;32;) (type 15)))
  (import "env" "mBufferStorageLoadFromAddress" (func (;33;) (type 1)))
  (import "env" "bigIntSign" (func (;34;) (type 7)))
  (import "env" "bigIntCmp" (func (;35;) (type 4)))
  (import "env" "managedWriteLog" (func (;36;) (type 2)))
  (import "env" "managedMultiTransferDCDTNFTExecute" (func (;37;) (type 16)))
  (import "env" "bigIntFinishUnsigned" (func (;38;) (type 6)))
  (import "env" "mBufferGetBytes" (func (;39;) (type 4)))
  (import "env" "mBufferGetByteSlice" (func (;40;) (type 11)))
  (func (;41;) (type 7) (param i32) (result i32)
    (local i32)
    call 42
    local.tee 1
    i64.const 0
    call 0
    local.get 1
    local.get 1
    local.get 0
    call 1
    local.get 1)
  (func (;42;) (type 3) (result i32)
    (local i32)
    i32.const 0
    i32.const 0
    i32.load offset=132068
    i32.const -1
    i32.add
    local.tee 0
    i32.store offset=132068
    local.get 0)
  (func (;43;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 2
    unreachable)
  (func (;44;) (type 2) (param i32 i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.load offset=4
          local.tee 3
          local.get 1
          i32.load
          local.tee 4
          i32.load offset=88
          local.tee 5
          i32.gt_u
          br_if 0 (;@3;)
          local.get 1
          local.get 3
          i32.const 1
          i32.add
          i32.store offset=4
          local.get 3
          i32.const 1
          i32.ne
          br_if 1 (;@2;)
          local.get 2
          i32.const 8
          i32.add
          local.get 4
          call 45
          call 46
          local.get 0
          local.get 2
          i32.const 8
          i32.add
          i32.const 40
          call 204
          drop
          br 2 (;@1;)
        end
        local.get 0
        i32.const 2
        i32.store8 offset=32
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 3
        local.get 5
        i32.eq
        br_if 0 (;@2;)
        local.get 2
        i32.const 8
        i32.add
        local.get 4
        i32.load
        local.get 3
        call 47
        local.get 0
        local.get 2
        i32.const 8
        i32.add
        i32.const 40
        call 204
        drop
        br 1 (;@1;)
      end
      local.get 2
      i32.const 8
      i32.add
      local.get 4
      call 48
      call 46
      local.get 0
      local.get 2
      i32.const 8
      i32.add
      i32.const 40
      call 204
      drop
    end
    local.get 2
    i32.const 48
    i32.add
    global.set 0)
  (func (;45;) (type 7) (param i32) (result i32)
    block  ;; label = @1
      local.get 0
      i32.load offset=88
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const 8
      i32.add
      return
    end
    i32.const 132056
    i32.const 9
    call 2
    unreachable)
  (func (;46;) (type 2) (param i32 i32)
    (local i64 i64 i32 i32 i32)
    local.get 1
    i64.load
    local.set 2
    local.get 1
    i64.load offset=8
    local.set 3
    local.get 1
    i32.load offset=16
    call 49
    local.set 4
    local.get 1
    i32.load offset=20
    call 49
    local.set 5
    local.get 1
    i32.load offset=24
    call 49
    local.set 6
    local.get 0
    local.get 1
    i32.load offset=28
    call 41
    i32.store offset=28
    local.get 0
    local.get 6
    i32.store offset=24
    local.get 0
    local.get 5
    i32.store offset=20
    local.get 0
    local.get 4
    i32.store offset=16
    local.get 0
    local.get 3
    i64.store offset=8
    local.get 0
    local.get 2
    i64.store
    local.get 0
    local.get 1
    i32.load8_u offset=32
    i32.store8 offset=32)
  (func (;47;) (type 1) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 131
    call 132)
  (func (;48;) (type 7) (param i32) (result i32)
    block  ;; label = @1
      local.get 0
      i32.load offset=88
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const 48
      i32.add
      return
    end
    i32.const 132056
    i32.const 9
    call 2
    unreachable)
  (func (;49;) (type 7) (param i32) (result i32)
    (local i32)
    call 3
    local.tee 1
    local.get 0
    call 4
    drop
    local.get 1)
  (func (;50;) (type 7) (param i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 1
    global.set 0
    call 51
    local.set 2
    local.get 1
    local.get 0
    i32.load
    call 52
    i32.store offset=28
    local.get 1
    i32.const 0
    i32.store offset=24
    local.get 1
    local.get 0
    i32.store offset=20
    loop (result i32)  ;; label = @1
      local.get 1
      i32.const 8
      i32.add
      local.get 1
      i32.const 20
      i32.add
      call 53
      block  ;; label = @2
        local.get 1
        i32.load offset=8
        br_if 0 (;@2;)
        local.get 1
        i32.const 32
        i32.add
        global.set 0
        local.get 2
        return
      end
      local.get 2
      local.get 1
      i32.load offset=12
      call 49
      call 54
      br 0 (;@1;)
    end)
  (func (;51;) (type 3) (result i32)
    (local i32)
    call 42
    local.tee 0
    i32.const 131988
    i32.const 0
    call 21
    drop
    local.get 0)
  (func (;52;) (type 7) (param i32) (result i32)
    local.get 0
    call 15)
  (func (;53;) (type 2) (param i32 i32)
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
      call 92
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
  (func (;54;) (type 2) (param i32 i32)
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
    call 13
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;55;) (type 8) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call 56
    call 57)
  (func (;56;) (type 8) (param i32 i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 3
    i32.const 12
    i32.add
    i32.const 4
    local.get 1
    local.get 2
    call 105
    local.get 0
    local.get 3
    i32.load offset=12
    local.tee 4
    i32.const 24
    i32.shl
    local.get 4
    i32.const 65280
    i32.and
    i32.const 8
    i32.shl
    i32.or
    local.get 4
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 4
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    local.get 1
    local.get 2
    call 106
    local.set 4
    local.get 3
    i32.const 16
    i32.add
    global.set 0
    local.get 4)
  (func (;57;) (type 7) (param i32) (result i32)
    (local i32)
    local.get 0
    call 42
    local.tee 1
    call 24
    drop
    local.get 1)
  (func (;58;) (type 2) (param i32 i32)
    local.get 0
    call 59
    local.get 1
    call 60)
  (func (;59;) (type 7) (param i32) (result i32)
    (local i32)
    call 42
    local.tee 1
    local.get 0
    call 23
    drop
    local.get 1)
  (func (;60;) (type 2) (param i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    call 52
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
    call 111
    local.get 1
    local.get 0
    call 114
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;61;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 43
    unreachable)
  (func (;62;) (type 17) (param i64 i32 i32 i32 i32) (result i32)
    (local i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    local.get 4
    call 42
    local.tee 5
    call 5
    drop
    local.get 5)
  (func (;63;) (type 3) (result i32)
    (local i32)
    call 42
    local.tee 0
    call 6
    local.get 0)
  (func (;64;) (type 7) (param i32) (result i32)
    local.get 0
    call 65
    i32.const 142084
    call 7
    i32.const 0
    i32.gt_s)
  (func (;65;) (type 6) (param i32)
    local.get 0
    i32.const 142084
    call 39
    drop)
  (func (;66;) (type 7) (param i32) (result i32)
    local.get 0
    call 65
    i32.const 142084
    call 8)
  (func (;67;) (type 15)
    (local i32 i32)
    call 42
    local.tee 0
    call 9
    call 42
    local.tee 1
    call 10
    block  ;; label = @1
      local.get 0
      local.get 1
      call 11
      i32.const 1
      i32.lt_s
      br_if 0 (;@1;)
      return
    end
    i32.const 131988
    i32.const 36
    call 2
    unreachable)
  (func (;68;) (type 7) (param i32) (result i32)
    (local i32)
    local.get 0
    call 42
    local.tee 1
    call 12
    drop
    local.get 1)
  (func (;69;) (type 18) (param i32 i32 i32 i32)
    (local i32)
    i32.const 131118
    i32.const 23
    call 70
    local.tee 4
    local.get 0
    local.get 1
    call 13
    drop
    local.get 4
    i32.const 131141
    i32.const 3
    call 13
    drop
    local.get 4
    local.get 2
    local.get 3
    call 13
    drop
    local.get 4
    call 14
    unreachable)
  (func (;70;) (type 4) (param i32 i32) (result i32)
    (local i32)
    call 42
    local.tee 2
    local.get 0
    local.get 1
    call 21
    drop
    local.get 2)
  (func (;71;) (type 7) (param i32) (result i32)
    (local i32)
    call 51
    local.set 1
    block  ;; label = @1
      loop  ;; label = @2
        local.get 0
        i32.load
        i32.const 0
        i32.load offset=142080
        i32.ge_s
        br_if 1 (;@1;)
        local.get 1
        local.get 0
        i32.const 131494
        i32.const 9
        call 72
        call 68
        call 54
        br 0 (;@2;)
      end
    end
    local.get 1)
  (func (;72;) (type 8) (param i32 i32 i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 3
      i32.const 0
      i32.load offset=142080
      i32.lt_s
      br_if 0 (;@1;)
      local.get 1
      local.get 2
      i32.const 131144
      i32.const 17
      call 69
      unreachable
    end
    local.get 0
    local.get 3
    i32.const 1
    i32.add
    i32.store
    local.get 3)
  (func (;73;) (type 2) (param i32 i32)
    (local i32)
    i32.const 0
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.load
        i32.const 0
        i32.load offset=142080
        i32.lt_s
        br_if 0 (;@2;)
        i32.const 1
        local.set 2
        br 1 (;@1;)
      end
      local.get 1
      i32.const 131503
      i32.const 15
      call 72
      call 68
      local.tee 1
      call 15
      i32.const 32
      i32.eq
      br_if 0 (;@1;)
      i32.const 131503
      i32.const 15
      i32.const 131342
      i32.const 16
      call 69
      unreachable
    end
    local.get 0
    local.get 1
    i32.store offset=4
    local.get 0
    local.get 2
    i32.store)
  (func (;74;) (type 3) (result i32)
    i32.const 0
    call 68)
  (func (;75;) (type 3) (result i32)
    (local i32)
    i32.const 1
    call 42
    local.tee 0
    call 16
    local.get 0)
  (func (;76;) (type 3) (result i32)
    (local i64)
    block  ;; label = @1
      i32.const 0
      call 17
      local.tee 0
      i64.const 4294967296
      i64.lt_u
      br_if 0 (;@1;)
      i32.const 131442
      i32.const 21
      i32.const 131072
      i32.const 14
      call 69
      unreachable
    end
    local.get 0
    i32.wrap_i64)
  (func (;77;) (type 12) (result i64)
    i32.const 0
    call 17)
  (func (;78;) (type 6) (param i32)
    block  ;; label = @1
      i32.const 0
      i32.load offset=142080
      local.get 0
      i32.gt_s
      br_if 0 (;@1;)
      return
    end
    i32.const 131161
    i32.const 18
    call 2
    unreachable)
  (func (;79;) (type 6) (param i32)
    block  ;; label = @1
      call 18
      local.get 0
      i32.ne
      br_if 0 (;@1;)
      return
    end
    i32.const 131179
    i32.const 25
    call 2
    unreachable)
  (func (;80;) (type 6) (param i32)
    block  ;; label = @1
      i32.const 0
      i32.load offset=142080
      local.get 0
      i32.lt_s
      br_if 0 (;@1;)
      return
    end
    i32.const 131144
    i32.const 17
    call 2
    unreachable)
  (func (;81;) (type 15)
    i32.const 0
    call 18
    i32.store offset=142080)
  (func (;82;) (type 6) (param i32)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 1
    global.set 0
    block  ;; label = @1
      local.get 0
      i64.load
      i64.const 0
      i64.ne
      br_if 0 (;@1;)
      local.get 0
      i64.load offset=8
      call 19
      local.get 1
      local.get 0
      i32.load offset=16
      call 52
      i32.store offset=28
      local.get 1
      i32.const 0
      i32.store offset=24
      local.get 1
      local.get 0
      i32.const 16
      i32.add
      i32.store offset=20
      loop  ;; label = @2
        local.get 1
        i32.const 8
        i32.add
        local.get 1
        i32.const 20
        i32.add
        call 53
        local.get 1
        i32.load offset=8
        i32.eqz
        br_if 1 (;@1;)
        local.get 1
        i32.load offset=12
        call 20
        drop
        br 0 (;@2;)
      end
    end
    local.get 1
    i32.const 32
    i32.add
    global.set 0)
  (func (;83;) (type 6) (param i32)
    local.get 0
    call 84
    call 20
    drop)
  (func (;84;) (type 7) (param i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      call 123
      local.tee 1
      call 15
      i32.const 32
      i32.eq
      br_if 0 (;@1;)
      local.get 0
      i32.const 131342
      i32.const 16
      call 110
      unreachable
    end
    local.get 1)
  (func (;85;) (type 6) (param i32)
    local.get 0
    call 86
    call 19)
  (func (;86;) (type 9) (param i32) (result i64)
    local.get 0
    local.get 0
    call 141)
  (func (;87;) (type 19) (param i32 i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32 i32 i64 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 5
    global.set 0
    local.get 5
    i32.const 16
    i32.add
    i32.const 131204
    i32.const 20
    call 88
    local.get 5
    i32.load offset=16
    local.set 6
    local.get 5
    i32.load offset=20
    local.set 7
    call 51
    drop
    local.get 7
    local.get 3
    call 49
    call 54
    local.get 4
    call 89
    local.set 3
    call 51
    local.tee 8
    local.get 3
    i64.extend_i32_u
    call 90
    local.get 7
    local.get 8
    call 54
    local.get 4
    call 52
    local.set 9
    local.get 5
    i32.const 32
    i32.add
    local.set 10
    i32.const 0
    local.set 3
    loop  ;; label = @1
      block  ;; label = @2
        local.get 3
        i32.const 16
        i32.add
        local.tee 8
        local.get 9
        i32.le_u
        br_if 0 (;@2;)
        local.get 5
        local.get 6
        local.get 7
        local.get 1
        local.get 2
        call 91
        local.get 5
        i32.load offset=4
        local.set 7
        local.get 0
        local.get 5
        i32.load
        i32.store
        local.get 0
        local.get 7
        i32.store offset=4
        local.get 5
        i32.const 48
        i32.add
        global.set 0
        return
      end
      local.get 10
      i64.const 0
      i64.store
      local.get 5
      i64.const 0
      i64.store offset=24
      local.get 4
      local.get 3
      local.get 5
      i32.const 24
      i32.add
      i32.const 16
      call 92
      drop
      local.get 5
      i32.const 0
      i32.store offset=44
      local.get 5
      i32.const 24
      i32.add
      local.get 5
      i32.const 44
      i32.add
      call 93
      local.set 3
      local.get 5
      i32.const 24
      i32.add
      local.get 5
      i32.const 44
      i32.add
      call 94
      local.set 11
      local.get 5
      i32.const 24
      i32.add
      local.get 5
      i32.const 44
      i32.add
      call 93
      local.set 12
      local.get 7
      local.get 3
      call 95
      local.get 5
      i32.const 8
      i32.add
      local.get 6
      local.get 7
      local.get 11
      call 96
      local.get 5
      i32.load offset=8
      local.set 6
      local.get 5
      i32.load offset=12
      local.tee 7
      local.get 12
      call 97
      local.get 8
      local.set 3
      br 0 (;@1;)
    end)
  (func (;88;) (type 1) (param i32 i32 i32)
    local.get 1
    local.get 2
    call 70
    local.set 2
    local.get 0
    call 51
    i32.store offset=4
    local.get 0
    local.get 2
    i32.store)
  (func (;89;) (type 7) (param i32) (result i32)
    local.get 0
    call 52
    i32.const 4
    i32.shr_u)
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
    local.get 2
    i32.const 8
    i32.add
    call 162
    local.get 0
    local.get 2
    i32.load
    local.get 2
    i32.load offset=4
    call 21
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;91;) (type 19) (param i32 i32 i32 i32 i32)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 5
    global.set 0
    block  ;; label = @1
      local.get 3
      call 52
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      local.get 3
      call 95
      local.get 4
      call 52
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
        call 92
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
        call 95
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
  (func (;92;) (type 11) (param i32 i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    call 113)
  (func (;93;) (type 4) (param i32 i32) (result i32)
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
      call 119
      unreachable
    end
    local.get 3
    i32.const 16
    call 119
    unreachable)
  (func (;94;) (type 20) (param i32 i32) (result i64)
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
      call 119
      unreachable
    end
    local.get 3
    i32.const 16
    call 119
    unreachable)
  (func (;95;) (type 2) (param i32 i32)
    call 51
    drop
    local.get 0
    local.get 1
    call 49
    call 54)
  (func (;96;) (type 21) (param i32 i32 i32 i64)
    local.get 2
    local.get 3
    call 102
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;97;) (type 2) (param i32 i32)
    call 51
    drop
    local.get 0
    local.get 1
    call 59
    call 54)
  (func (;98;) (type 19) (param i32 i32 i32 i32 i32)
    (local i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 5
    global.set 0
    local.get 5
    i32.const 40
    i32.add
    i32.const 131224
    i32.const 15
    call 88
    local.get 5
    i32.const 32
    i32.add
    local.get 5
    i32.load offset=40
    local.get 5
    i32.load offset=44
    local.get 4
    i32.load offset=8
    call 99
    local.get 5
    i32.const 24
    i32.add
    local.get 5
    i32.load offset=32
    local.get 5
    i32.load offset=36
    local.get 4
    i64.load
    call 96
    local.get 5
    i32.const 16
    i32.add
    local.get 5
    i32.load offset=24
    local.get 5
    i32.load offset=28
    local.get 4
    i32.load offset=12
    call 100
    local.get 5
    i32.load offset=16
    local.set 4
    local.get 5
    i32.load offset=20
    local.tee 6
    local.get 3
    call 95
    local.get 5
    i32.const 8
    i32.add
    local.get 4
    local.get 6
    local.get 1
    local.get 2
    call 91
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
  (func (;99;) (type 18) (param i32 i32 i32 i32)
    call 51
    drop
    local.get 2
    local.get 3
    i32.load
    call 49
    call 54
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;100;) (type 18) (param i32 i32 i32 i32)
    call 51
    drop
    local.get 2
    local.get 3
    i32.load
    call 59
    call 54
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;101;) (type 19) (param i32 i32 i32 i32 i32)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 5
    global.set 0
    local.get 5
    i32.const 24
    i32.add
    i32.const 131239
    i32.const 12
    call 88
    local.get 5
    i32.const 16
    i32.add
    local.get 5
    i32.load offset=24
    local.get 5
    i32.load offset=28
    local.get 3
    call 99
    local.get 5
    i32.const 8
    i32.add
    local.get 5
    i32.load offset=16
    local.get 5
    i32.load offset=20
    local.get 4
    call 100
    local.get 5
    local.get 5
    i32.load offset=8
    local.get 5
    i32.load offset=12
    local.get 1
    local.get 2
    call 91
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
  (func (;102;) (type 0) (param i32 i64)
    (local i32)
    call 51
    local.tee 2
    local.get 1
    call 90
    local.get 0
    local.get 2
    call 54)
  (func (;103;) (type 17) (param i64 i32 i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    local.get 4
    call 62)
  (func (;104;) (type 22) (param i64 i32 i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    local.get 4
    call 62
    drop)
  (func (;105;) (type 19) (param i32 i32 i32 i32 i32)
    block  ;; label = @1
      local.get 0
      i32.const 8
      i32.add
      local.get 0
      i32.load
      local.get 1
      local.get 2
      call 120
      i32.eqz
      br_if 0 (;@1;)
      local.get 3
      local.get 4
      call 157
      unreachable
    end
    local.get 0
    local.get 0
    i32.load
    local.get 2
    i32.add
    i32.store)
  (func (;106;) (type 11) (param i32 i32 i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 4
    global.set 0
    local.get 4
    i32.const 8
    i32.add
    local.get 0
    i32.load offset=8
    local.get 0
    i32.load
    local.tee 5
    local.get 1
    call 109
    block  ;; label = @1
      local.get 4
      i32.load offset=8
      i32.const 1
      i32.ne
      br_if 0 (;@1;)
      local.get 4
      i32.load offset=12
      local.set 3
      local.get 0
      local.get 5
      local.get 1
      i32.add
      i32.store
      local.get 4
      i32.const 16
      i32.add
      global.set 0
      local.get 3
      return
    end
    local.get 2
    local.get 3
    i32.const 131262
    i32.const 15
    call 69
    unreachable)
  (func (;107;) (type 4) (param i32 i32) (result i32)
    (local i32 i32 i32)
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
    call 108
    local.get 2
    local.get 0
    i32.load offset=8
    local.get 0
    i32.load
    local.tee 3
    local.get 2
    i32.load offset=12
    local.tee 4
    i32.const 24
    i32.shl
    local.get 4
    i32.const 65280
    i32.and
    i32.const 8
    i32.shl
    i32.or
    local.get 4
    i32.const 8
    i32.shr_u
    i32.const 65280
    i32.and
    local.get 4
    i32.const 24
    i32.shr_u
    i32.or
    i32.or
    local.tee 4
    call 109
    block  ;; label = @1
      local.get 2
      i32.load
      i32.const 1
      i32.eq
      br_if 0 (;@1;)
      local.get 1
      i32.const 131262
      i32.const 15
      call 110
      unreachable
    end
    local.get 2
    i32.load offset=4
    local.set 1
    local.get 0
    local.get 3
    local.get 4
    i32.add
    i32.store
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 1)
  (func (;108;) (type 18) (param i32 i32 i32 i32)
    block  ;; label = @1
      local.get 0
      i32.const 8
      i32.add
      local.get 0
      i32.load
      local.get 1
      local.get 2
      call 120
      i32.eqz
      br_if 0 (;@1;)
      local.get 3
      call 158
      unreachable
    end
    local.get 0
    local.get 0
    i32.load
    local.get 2
    i32.add
    i32.store)
  (func (;109;) (type 18) (param i32 i32 i32 i32)
    (local i32)
    local.get 1
    local.get 2
    local.get 3
    call 3
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
  (func (;110;) (type 1) (param i32 i32 i32)
    (local i32)
    i32.const 131311
    i32.const 27
    call 70
    local.tee 3
    local.get 0
    call 4
    drop
    local.get 3
    i32.const 131141
    i32.const 3
    call 13
    drop
    local.get 3
    local.get 1
    local.get 2
    call 13
    drop
    local.get 3
    call 14
    unreachable)
  (func (;111;) (type 1) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call 13
    drop)
  (func (;112;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    i64.load
    call 102
    local.get 0
    local.get 1
    i64.load offset=8
    call 102
    local.get 0
    local.get 1
    i32.load offset=16
    call 95
    local.get 0
    local.get 1
    i32.load offset=20
    call 95
    local.get 0
    local.get 1
    i32.load offset=24
    call 95
    local.get 0
    local.get 1
    i32.load offset=28
    call 97)
  (func (;113;) (type 11) (param i32 i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 3
    local.get 2
    call 40
    i32.const 0
    i32.ne)
  (func (;114;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 4
    drop)
  (func (;115;) (type 2) (param i32 i32)
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
    call 92
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
    call 93
    local.set 3
    local.get 2
    i32.const 8
    i32.add
    local.get 2
    i32.const 28
    i32.add
    call 94
    local.set 4
    local.get 2
    i32.const 8
    i32.add
    local.get 2
    i32.const 28
    i32.add
    call 93
    local.set 5
    block  ;; label = @1
      local.get 1
      i32.eqz
      br_if 0 (;@1;)
      i32.const 131277
      i32.const 29
      call 2
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
  (func (;116;) (type 2) (param i32 i32)
    (local i32 i32 i64)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    i32.load8_u offset=32
    i32.store8 offset=40
    local.get 2
    local.get 1
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
    i32.store offset=24
    local.get 2
    local.get 1
    i32.load offset=20
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
    i32.store offset=28
    local.get 2
    local.get 1
    i32.load offset=24
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
    i32.store offset=32
    local.get 2
    local.get 1
    i32.load offset=28
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
    i32.store offset=36
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
    i64.store offset=8
    local.get 2
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
    i64.store offset=16
    local.get 0
    local.get 2
    i32.const 8
    i32.add
    i32.const 33
    call 13
    drop
    local.get 2
    i32.const 48
    i32.add
    global.set 0)
  (func (;117;) (type 2) (param i32 i32)
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
    call 13
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;118;) (type 0) (param i32 i64)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
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
    i64.const 65280
    i64.and
    local.get 1
    i64.const 56
    i64.shr_u
    i64.or
    i64.or
    i64.or
    i64.store offset=8
    local.get 0
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    call 13
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;119;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 197
    unreachable)
  (func (;120;) (type 11) (param i32 i32 i32 i32) (result i32)
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
              call 52
              local.tee 5
              i32.const 10000
              i32.gt_u
              br_if 1 (;@4;)
              i32.const 0
              i32.load8_u offset=142076
              i32.const 255
              i32.and
              br_if 1 (;@4;)
              i32.const 0
              local.get 5
              i32.store offset=142072
              i32.const 0
              i32.const 1
              i32.store8 offset=142076
              local.get 4
              i32.const 0
              i32.const 132072
              local.get 5
              call 92
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
            i32.load offset=142072
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
            i32.const 132072
            i32.add
            local.get 3
            call 121
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
          call 92
          local.set 4
        end
        local.get 4
        return
      end
      local.get 1
      local.get 0
      call 119
      unreachable
    end
    local.get 0
    i32.const 10000
    call 119
    unreachable)
  (func (;121;) (type 18) (param i32 i32 i32 i32)
    block  ;; label = @1
      local.get 1
      local.get 3
      i32.eq
      br_if 0 (;@1;)
      local.get 1
      local.get 3
      call 198
      unreachable
    end
    local.get 0
    local.get 2
    local.get 1
    call 204
    drop)
  (func (;122;) (type 3) (result i32)
    i32.const -10
    i64.const 0
    call 0
    i32.const -10)
  (func (;123;) (type 7) (param i32) (result i32)
    (local i32)
    local.get 0
    call 42
    local.tee 1
    call 25
    drop
    local.get 1)
  (func (;124;) (type 2) (param i32 i32)
    (local i32)
    i32.const 131988
    i32.const 0
    call 70
    local.set 2
    local.get 1
    i64.load
    local.get 2
    call 125
    local.get 1
    i64.load offset=8
    local.get 2
    call 125
    local.get 1
    i32.load offset=16
    local.get 2
    call 60
    local.get 1
    i32.load offset=20
    local.get 2
    call 60
    local.get 1
    i32.load offset=24
    local.get 2
    call 60
    local.get 1
    i32.load offset=28
    call 59
    local.get 2
    call 60
    local.get 1
    i32.load8_u offset=32
    local.get 2
    call 126
    local.get 0
    local.get 2
    call 26
    drop)
  (func (;125;) (type 23) (param i64 i32)
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
    call 111
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;126;) (type 2) (param i32 i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 0
    i32.store8 offset=15
    local.get 1
    local.get 2
    i32.const 15
    i32.add
    i32.const 1
    call 111
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;127;) (type 6) (param i32)
    i32.const -20
    i32.const 131988
    i32.const 0
    call 21
    drop
    local.get 0
    i32.const -20
    call 26
    drop)
  (func (;128;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 26
    drop)
  (func (;129;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    i64.extend_i32_u
    call 130)
  (func (;130;) (type 0) (param i32 i64)
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
    local.get 2
    i32.const 8
    i32.add
    call 162
    local.get 0
    local.get 2
    i32.load
    local.get 2
    i32.load offset=4
    call 70
    call 26
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;131;) (type 4) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    call 49
    local.tee 0
    i32.const 131306
    i32.const 5
    call 13
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
    call 13
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func (;132;) (type 2) (param i32 i32)
    (local i32 i32 i32 i32 i64 i64 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    i32.const 0
    local.set 3
    local.get 1
    call 123
    local.tee 4
    call 15
    local.set 5
    local.get 2
    i32.const 0
    i32.store8 offset=24
    local.get 2
    local.get 5
    i32.store offset=20
    local.get 2
    local.get 4
    i32.store offset=16
    local.get 2
    local.get 5
    i32.store offset=12
    local.get 2
    i32.const 0
    i32.store offset=8
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    call 144
    local.set 6
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    call 144
    local.set 7
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    call 107
    local.set 5
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    call 107
    local.set 4
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    call 107
    local.set 8
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    call 107
    call 57
    local.set 9
    local.get 2
    i32.const 0
    i32.store8 offset=31
    local.get 2
    i32.const 8
    i32.add
    local.get 2
    i32.const 31
    i32.add
    i32.const 1
    local.get 1
    call 108
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          i32.load8_u offset=31
          br_table 2 (;@1;) 1 (;@2;) 0 (;@3;)
        end
        local.get 1
        i32.const 131661
        i32.const 13
        call 110
        unreachable
      end
      i32.const 1
      local.set 3
    end
    block  ;; label = @1
      local.get 2
      i32.load offset=12
      local.get 2
      i32.load offset=8
      i32.ne
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 2
        i32.load8_u offset=24
        i32.eqz
        br_if 0 (;@2;)
        i32.const 0
        i32.const 0
        i32.store offset=142072
        i32.const 0
        i32.const 0
        i32.store8 offset=142076
      end
      local.get 0
      local.get 3
      i32.store8 offset=32
      local.get 0
      local.get 9
      i32.store offset=28
      local.get 0
      local.get 8
      i32.store offset=24
      local.get 0
      local.get 4
      i32.store offset=20
      local.get 0
      local.get 5
      i32.store offset=16
      local.get 0
      local.get 7
      i64.store offset=8
      local.get 0
      local.get 6
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
    call 110
    unreachable)
  (func (;133;) (type 18) (param i32 i32 i32 i32)
    block  ;; label = @1
      local.get 2
      call 134
      local.get 3
      i32.ge_u
      br_if 0 (;@1;)
      i32.const 132024
      i32.const 18
      call 2
      unreachable
    end
    local.get 0
    local.get 1
    local.get 3
    call 47)
  (func (;134;) (type 7) (param i32) (result i32)
    local.get 0
    call 135)
  (func (;135;) (type 7) (param i32) (result i32)
    (local i64)
    block  ;; label = @1
      local.get 0
      local.get 0
      call 141
      local.tee 1
      i64.const 4294967296
      i64.lt_u
      br_if 0 (;@1;)
      local.get 0
      i32.const 131072
      i32.const 14
      call 110
      unreachable
    end
    local.get 1
    i32.wrap_i64)
  (func (;136;) (type 0) (param i32 i64)
    block  ;; label = @1
      local.get 0
      call 137
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      call 138
    end)
  (func (;137;) (type 7) (param i32) (result i32)
    local.get 0
    i32.const -25
    call 25
    drop
    i32.const -25
    call 15)
  (func (;138;) (type 0) (param i32 i64)
    local.get 0
    local.get 1
    call 130)
  (func (;139;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    call 128)
  (func (;140;) (type 7) (param i32) (result i32)
    local.get 0
    call 123
    call 57)
  (func (;141;) (type 20) (param i32 i32) (result i64)
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
      call 123
      local.tee 3
      call 15
      local.tee 0
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 1
      i32.const 131072
      i32.const 14
      call 110
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
    call 113
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
  (func (;142;) (type 0) (param i32 i64)
    local.get 0
    local.get 1
    call 102)
  (func (;143;) (type 4) (param i32 i32) (result i32)
    (local i32)
    call 51
    local.tee 2
    local.get 0
    local.get 1
    call 70
    call 54
    local.get 2)
  (func (;144;) (type 20) (param i32 i32) (result i64)
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
    call 108
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
  (func (;145;) (type 2) (param i32 i32)
    (local i32 i32 i32 i32 i32 i64 i64 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    i32.const 2
    local.set 3
    block  ;; label = @1
      local.get 1
      i32.load offset=4
      local.tee 4
      i32.const 33
      i32.add
      local.tee 5
      local.get 1
      i32.load offset=8
      i32.gt_u
      br_if 0 (;@1;)
      local.get 1
      i32.load
      local.set 3
      local.get 2
      i32.const 8
      i32.add
      i32.const 0
      i32.const 33
      call 203
      drop
      local.get 3
      i32.load
      local.get 4
      local.get 2
      i32.const 8
      i32.add
      i32.const 33
      call 92
      drop
      local.get 2
      i32.load8_u offset=40
      local.set 6
      local.get 2
      i64.load offset=8
      local.set 7
      local.get 2
      i64.load offset=16
      local.set 8
      local.get 2
      i32.load offset=24
      local.set 3
      local.get 2
      i32.load offset=28
      local.set 4
      local.get 2
      i32.load offset=32
      local.set 9
      local.get 2
      i32.load offset=36
      local.set 10
      local.get 1
      local.get 5
      i32.store offset=4
      local.get 0
      local.get 10
      i32.const 24
      i32.shl
      local.get 10
      i32.const 65280
      i32.and
      i32.const 8
      i32.shl
      i32.or
      local.get 10
      i32.const 8
      i32.shr_u
      i32.const 65280
      i32.and
      local.get 10
      i32.const 24
      i32.shr_u
      i32.or
      i32.or
      i32.store offset=28
      local.get 0
      local.get 9
      i32.const 24
      i32.shl
      local.get 9
      i32.const 65280
      i32.and
      i32.const 8
      i32.shl
      i32.or
      local.get 9
      i32.const 8
      i32.shr_u
      i32.const 65280
      i32.and
      local.get 9
      i32.const 24
      i32.shr_u
      i32.or
      i32.or
      i32.store offset=24
      local.get 0
      local.get 4
      i32.const 24
      i32.shl
      local.get 4
      i32.const 65280
      i32.and
      i32.const 8
      i32.shl
      i32.or
      local.get 4
      i32.const 8
      i32.shr_u
      i32.const 65280
      i32.and
      local.get 4
      i32.const 24
      i32.shr_u
      i32.or
      i32.or
      i32.store offset=20
      local.get 0
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
      i32.store offset=16
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
      i64.store offset=8
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
      local.get 6
      i32.const 0
      i32.ne
      local.set 3
    end
    local.get 0
    local.get 3
    i32.store8 offset=32
    local.get 2
    i32.const 48
    i32.add
    global.set 0)
  (func (;146;) (type 24) (param i32 i64 i64) (result i32)
    (local i32 i32)
    block  ;; label = @1
      local.get 0
      i32.load offset=88
      local.tee 3
      i32.eqz
      br_if 0 (;@1;)
      i32.const 1
      local.set 4
      block  ;; label = @2
        local.get 3
        call 147
        call 135
        i32.ge_u
        br_if 0 (;@2;)
        local.get 1
        local.get 2
        i64.gt_u
        br_if 1 (;@1;)
        i32.const 0
        local.set 4
        call 27
        local.tee 2
        local.get 0
        call 45
        i64.load
        local.tee 1
        i64.lt_u
        br_if 0 (;@2;)
        local.get 2
        local.get 1
        i64.sub
        call 148
        call 86
        i64.ge_u
        local.set 4
      end
      local.get 4
      return
    end
    i32.const 0)
  (func (;147;) (type 3) (result i32)
    i32.const 131729
    i32.const 14
    call 70)
  (func (;148;) (type 3) (result i32)
    i32.const 131743
    i32.const 23
    call 70)
  (func (;149;) (type 6) (param i32)
    (local i32 i32 i64 i64 i32)
    global.get 0
    i32.const 256
    i32.sub
    local.tee 1
    global.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.load
        local.tee 2
        call 52
        i32.eqz
        br_if 0 (;@2;)
        call 150
        call 86
        local.set 3
        local.get 1
        i32.const 8
        i32.add
        call 151
        call 86
        local.tee 4
        call 152
        call 51
        local.set 5
        local.get 1
        local.get 2
        call 52
        i32.store offset=116
        local.get 1
        i32.const 0
        i32.store offset=112
        local.get 1
        local.get 0
        i32.store offset=108
        loop  ;; label = @3
          local.get 1
          i32.const 120
          i32.add
          local.get 1
          i32.const 108
          i32.add
          call 145
          local.get 1
          i32.load8_u offset=152
          i32.const 2
          i32.eq
          br_if 2 (;@1;)
          block  ;; label = @4
            block  ;; label = @5
              local.get 1
              i32.const 8
              i32.add
              local.get 4
              local.get 3
              call 146
              br_if 0 (;@5;)
              local.get 1
              i32.const 8
              i32.add
              local.get 1
              i32.const 120
              i32.add
              call 153
              br 1 (;@4;)
            end
            local.get 1
            i32.const 160
            i32.add
            call 151
            call 86
            i64.const 1
            i64.add
            local.tee 4
            call 152
            local.get 1
            i32.const 160
            i32.add
            local.get 1
            i32.const 120
            i32.add
            call 153
            call 151
            local.get 4
            call 130
            local.get 1
            i32.const 8
            i32.add
            local.get 1
            i32.const 160
            i32.add
            i32.const 96
            call 204
            drop
          end
          local.get 5
          local.get 4
          call 118
          br 0 (;@3;)
        end
      end
      call 51
      drop
    end
    local.get 1
    i32.const 256
    i32.add
    global.set 0)
  (func (;150;) (type 3) (result i32)
    i32.const 131703
    i32.const 12
    call 70)
  (func (;151;) (type 3) (result i32)
    i32.const 131692
    i32.const 11
    call 70)
  (func (;152;) (type 0) (param i32 i64)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 128
    i32.sub
    local.tee 2
    global.set 0
    local.get 1
    i32.const 131715
    i32.const 14
    call 70
    local.tee 3
    call 125
    local.get 3
    call 49
    local.tee 4
    i32.const 131338
    i32.const 4
    call 13
    drop
    block  ;; label = @1
      block  ;; label = @2
        local.get 4
        call 135
        local.tee 5
        br_if 0 (;@2;)
        i32.const 2
        local.set 6
        i32.const 2
        local.set 7
        br 1 (;@1;)
      end
      local.get 2
      i32.const 8
      i32.add
      local.get 3
      local.get 4
      i32.const 1
      call 133
      local.get 2
      i32.const 56
      i32.add
      i32.const 24
      i32.add
      local.get 2
      i32.const 8
      i32.add
      i32.const 24
      i32.add
      local.tee 7
      i64.load
      i64.store
      local.get 2
      i32.const 56
      i32.add
      i32.const 16
      i32.add
      local.get 2
      i32.const 8
      i32.add
      i32.const 16
      i32.add
      local.tee 8
      i64.load
      i64.store
      local.get 2
      i32.const 56
      i32.add
      i32.const 8
      i32.add
      local.get 2
      i32.const 8
      i32.add
      i32.const 8
      i32.add
      local.tee 9
      i64.load
      i64.store
      local.get 2
      local.get 2
      i64.load offset=8
      i64.store offset=56
      local.get 2
      local.get 2
      i32.load offset=41 align=1
      i32.store offset=48
      local.get 2
      local.get 2
      i32.const 44
      i32.add
      local.tee 10
      i32.load align=1
      i32.store offset=51 align=1
      local.get 2
      i32.load8_u offset=40
      local.set 6
      local.get 2
      i32.const 8
      i32.add
      local.get 3
      local.get 4
      local.get 5
      call 133
      local.get 2
      i32.const 96
      i32.add
      i32.const 24
      i32.add
      local.get 7
      i64.load
      i64.store
      local.get 2
      i32.const 96
      i32.add
      i32.const 16
      i32.add
      local.get 8
      i64.load
      i64.store
      local.get 2
      i32.const 96
      i32.add
      i32.const 8
      i32.add
      local.get 9
      i64.load
      i64.store
      local.get 2
      local.get 2
      i64.load offset=8
      i64.store offset=96
      local.get 2
      local.get 2
      i32.load offset=41 align=1
      i32.store offset=88
      local.get 2
      local.get 10
      i32.load align=1
      i32.store offset=91 align=1
      local.get 2
      i32.load8_u offset=40
      local.set 7
    end
    local.get 0
    local.get 5
    i32.store offset=88
    local.get 0
    local.get 4
    i32.store offset=4
    local.get 0
    local.get 3
    i32.store
    local.get 0
    local.get 2
    i64.load offset=56
    i64.store offset=8
    local.get 0
    local.get 6
    i32.store8 offset=40
    local.get 0
    local.get 2
    i32.load offset=48
    i32.store offset=41 align=1
    local.get 0
    local.get 2
    i64.load offset=96
    i64.store offset=48
    local.get 0
    i32.const 44
    i32.add
    local.get 2
    i32.load offset=51 align=1
    i32.store align=1
    local.get 0
    i32.const 16
    i32.add
    local.get 2
    i32.const 56
    i32.add
    i32.const 8
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 24
    i32.add
    local.get 2
    i32.const 56
    i32.add
    i32.const 16
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 32
    i32.add
    local.get 2
    i32.const 56
    i32.add
    i32.const 24
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 56
    i32.add
    local.get 2
    i32.const 96
    i32.add
    i32.const 8
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 64
    i32.add
    local.get 2
    i32.const 96
    i32.add
    i32.const 16
    i32.add
    i64.load
    i64.store
    local.get 0
    i32.const 72
    i32.add
    local.get 2
    i32.const 96
    i32.add
    i32.const 24
    i32.add
    i64.load
    i64.store
    local.get 0
    local.get 7
    i32.store8 offset=80
    local.get 0
    local.get 2
    i32.load offset=88
    i32.store offset=81 align=1
    local.get 0
    i32.const 84
    i32.add
    local.get 2
    i32.load offset=91 align=1
    i32.store align=1
    local.get 2
    i32.const 128
    i32.add
    global.set 0)
  (func (;153;) (type 2) (param i32 i32)
    (local i32 i32 i32 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 2
    global.set 0
    block  ;; label = @1
      local.get 0
      i32.load offset=88
      local.tee 3
      br_if 0 (;@1;)
      local.get 2
      i32.const 8
      i32.add
      local.get 1
      call 46
      local.get 0
      i32.const 8
      i32.add
      local.get 2
      i32.const 8
      i32.add
      i32.const 40
      call 204
      drop
    end
    local.get 0
    i32.load offset=4
    local.tee 4
    call 134
    local.set 5
    local.get 0
    i32.load
    local.get 5
    i32.const 1
    i32.add
    local.tee 5
    call 131
    local.get 1
    call 124
    local.get 4
    local.get 5
    call 129
    local.get 0
    local.get 3
    i32.const 1
    i32.add
    i32.store offset=88
    local.get 0
    i32.const 48
    i32.add
    local.get 1
    i32.const 40
    call 204
    drop
    local.get 2
    i32.const 48
    i32.add
    global.set 0)
  (func (;154;) (type 6) (param i32)
    local.get 0
    call 150
    call 86
    call 155)
  (func (;155;) (type 0) (param i32 i64)
    (local i32 i32 i64 i64 i64)
    global.get 0
    i32.const 176
    i32.sub
    local.tee 2
    global.set 0
    local.get 2
    local.get 1
    call 152
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.load offset=88
        i32.eqz
        br_if 0 (;@2;)
        call 51
        local.set 3
        local.get 2
        i32.const 1
        i32.store offset=100
        local.get 2
        local.get 2
        i32.store offset=96
        loop  ;; label = @3
          local.get 2
          i32.const 104
          i32.add
          local.get 2
          i32.const 96
          i32.add
          call 44
          block  ;; label = @4
            local.get 2
            i32.load8_u offset=136
            i32.const 2
            i32.ne
            br_if 0 (;@4;)
            local.get 0
            local.get 3
            i32.store offset=16
            local.get 0
            local.get 1
            i64.store offset=8
            local.get 0
            i64.const 0
            i64.store
            br 3 (;@1;)
          end
          local.get 2
          i64.load offset=104
          local.set 4
          local.get 2
          i64.load offset=112
          local.set 5
          local.get 2
          i64.load offset=120
          local.set 6
          local.get 2
          local.get 2
          i64.load offset=128
          i64.store offset=168
          local.get 2
          local.get 6
          i64.store offset=160
          local.get 2
          local.get 5
          i64.store offset=152
          local.get 2
          local.get 4
          i64.store offset=144
          local.get 3
          local.get 2
          i32.const 144
          i32.add
          call 112
          br 0 (;@3;)
        end
      end
      local.get 0
      i64.const 1
      i64.store
    end
    local.get 2
    i32.const 176
    i32.add
    global.set 0)
  (func (;156;) (type 8) (param i32 i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 0
    i32.load offset=8
    local.set 4
    local.get 3
    i32.const 0
    i32.store offset=12
    block  ;; label = @1
      local.get 0
      i32.load
      local.get 4
      i32.const 2
      i32.shl
      local.get 3
      i32.const 12
      i32.add
      i32.const 4
      call 92
      br_if 0 (;@1;)
      local.get 3
      i32.load offset=12
      local.set 2
      local.get 0
      local.get 4
      i32.const 1
      i32.add
      i32.store offset=8
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
      call 49
      local.set 2
      local.get 3
      i32.const 16
      i32.add
      global.set 0
      local.get 2
      return
    end
    local.get 1
    local.get 2
    i32.const 131144
    i32.const 17
    call 69
    unreachable)
  (func (;157;) (type 2) (param i32 i32)
    local.get 0
    local.get 1
    i32.const 131262
    i32.const 15
    call 69
    unreachable)
  (func (;158;) (type 6) (param i32)
    local.get 0
    i32.const 131262
    i32.const 15
    call 110
    unreachable)
  (func (;159;) (type 8) (param i32 i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    i32.const 0
    local.set 4
    block  ;; label = @1
      local.get 1
      call 66
      local.get 0
      i32.ne
      br_if 0 (;@1;)
      local.get 1
      local.get 2
      i64.const 0
      i64.const 0
      call 28
      call 3
      local.tee 0
      call 3
      call 3
      call 3
      call 3
      local.tee 4
      i64.const 0
      call 28
      call 3
      call 29
      block  ;; label = @2
        local.get 4
        call 15
        br_if 0 (;@2;)
        local.get 4
        i32.const 131086
        i32.const 32
        call 21
        drop
      end
      local.get 3
      i32.const 0
      i32.store16 offset=14
      local.get 0
      i32.const 0
      local.get 3
      i32.const 14
      i32.add
      i32.const 2
      call 113
      drop
      local.get 3
      i32.load8_u offset=14
      i32.const 0
      i32.ne
      local.set 4
    end
    local.get 3
    i32.const 16
    i32.add
    global.set 0
    local.get 4)
  (func (;160;) (type 25) (param i32 i32 i32) (result i64)
    (local i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    i64.const 0
    i64.store offset=8
    local.get 0
    local.get 3
    i32.const 8
    i32.add
    i32.const 8
    local.get 1
    local.get 2
    call 105
    local.get 3
    i64.load offset=8
    local.set 4
    local.get 3
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
  (func (;161;) (type 25) (param i32 i32 i32) (result i64)
    (local i32 i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 3
    global.set 0
    local.get 3
    i64.const 0
    i64.store offset=8
    block  ;; label = @1
      local.get 0
      call 15
      local.tee 4
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 1
      local.get 2
      i32.const 131072
      i32.const 14
      call 69
      unreachable
    end
    local.get 0
    i32.const 0
    local.get 3
    i32.const 8
    i32.add
    local.get 4
    i32.sub
    i32.const 8
    i32.add
    local.get 4
    call 113
    drop
    local.get 3
    i64.load offset=8
    local.set 5
    local.get 3
    i32.const 16
    i32.add
    global.set 0
    local.get 5
    i64.const 56
    i64.shl
    local.get 5
    i64.const 65280
    i64.and
    i64.const 40
    i64.shl
    i64.or
    local.get 5
    i64.const 16711680
    i64.and
    i64.const 24
    i64.shl
    local.get 5
    i64.const 4278190080
    i64.and
    i64.const 8
    i64.shl
    i64.or
    i64.or
    local.get 5
    i64.const 8
    i64.shr_u
    i64.const 4278190080
    i64.and
    local.get 5
    i64.const 24
    i64.shr_u
    i64.const 16711680
    i64.and
    i64.or
    local.get 5
    i64.const 40
    i64.shr_u
    i64.const 65280
    i64.and
    local.get 5
    i64.const 56
    i64.shr_u
    i64.or
    i64.or
    i64.or)
  (func (;162;) (type 26) (param i32 i64 i32)
    (local i64 i64 i32 i32)
    local.get 2
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
    local.tee 3
    i64.const 65280
    i64.and
    local.get 1
    i64.const 56
    i64.shr_u
    local.tee 4
    i64.or
    i64.or
    i64.or
    i64.store align=1
    local.get 0
    i32.const 8
    i32.const 0
    local.get 4
    i32.wrap_i64
    i32.eqz
    local.tee 5
    local.get 1
    i64.const 48
    i64.shr_u
    i32.wrap_i64
    i32.const 255
    i32.and
    select
    local.tee 6
    local.get 5
    i32.add
    i32.const 0
    local.get 6
    local.get 3
    i32.wrap_i64
    i32.const 255
    i32.and
    select
    local.tee 5
    i32.add
    i32.const 0
    local.get 5
    local.get 1
    i64.const 32
    i64.shr_u
    i32.wrap_i64
    i32.const 255
    i32.and
    select
    local.tee 6
    i32.add
    i32.const 0
    local.get 6
    local.get 1
    i32.wrap_i64
    local.tee 5
    i32.const 24
    i32.shr_u
    select
    local.tee 6
    i32.add
    i32.const 0
    local.get 6
    local.get 5
    i32.const 16
    i32.shr_u
    i32.const 255
    i32.and
    select
    local.tee 6
    i32.add
    i32.const 0
    local.get 6
    local.get 5
    i32.const 8
    i32.shr_u
    i32.const 255
    i32.and
    select
    local.tee 5
    i32.add
    local.get 5
    i32.const 0
    local.get 1
    i64.eqz
    select
    i32.add
    local.tee 5
    i32.sub
    i32.store offset=4
    local.get 0
    local.get 2
    local.get 5
    i32.add
    i32.store)
  (func (;163;) (type 2) (param i32 i32)
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
    call 111
    local.get 2
    i32.const 16
    i32.add
    global.set 0)
  (func (;164;) (type 13) (param i64) (result i32)
    (local i32)
    local.get 0
    i32.const 131766
    i32.const 20
    call 70
    local.tee 1
    call 125
    local.get 1)
  (func (;165;) (type 3) (result i32)
    i32.const 131829
    i32.const 23
    call 70)
  (func (;166;) (type 3) (result i32)
    i32.const 131852
    i32.const 23
    call 70)
  (func (;167;) (type 3) (result i32)
    i32.const 131875
    i32.const 26
    call 70)
  (func (;168;) (type 7) (param i32) (result i32)
    (local i32)
    local.get 0
    i32.const 131971
    i32.const 16
    call 70
    local.tee 1
    call 60
    local.get 1)
  (func (;169;) (type 2) (param i32 i32)
    (local i32 i32 i32 i32 i32 i64)
    local.get 1
    i32.load offset=8
    call 49
    local.set 2
    local.get 1
    i32.load offset=12
    call 49
    local.set 3
    local.get 1
    i32.load offset=16
    call 49
    local.set 4
    local.get 1
    i32.load offset=20
    call 41
    local.set 5
    i32.const 2147483646
    local.set 6
    local.get 1
    i64.load
    local.set 7
    block  ;; label = @1
      local.get 1
      i32.load offset=24
      local.tee 1
      i32.const 2147483646
      i32.eq
      br_if 0 (;@1;)
      local.get 1
      call 49
      local.set 6
    end
    local.get 0
    local.get 5
    i32.store offset=20
    local.get 0
    local.get 4
    i32.store offset=16
    local.get 0
    local.get 3
    i32.store offset=12
    local.get 0
    local.get 2
    i32.store offset=8
    local.get 0
    local.get 6
    i32.store offset=24
    local.get 0
    local.get 7
    i64.store)
  (func (;170;) (type 7) (param i32) (result i32)
    local.get 0
    i32.const 131463
    i32.const 8
    call 156
    call 49)
  (func (;171;) (type 9) (param i32) (result i64)
    local.get 0
    i32.const 131463
    i32.const 8
    call 156
    i32.const 131463
    i32.const 8
    call 161)
  (func (;172;) (type 12) (result i64)
    (local i64)
    call 30
    local.tee 0
    i64.const -100000
    i64.add
    local.get 0
    local.get 0
    i64.const 100000
    i64.gt_u
    select)
  (func (;173;) (type 15)
    call 174
    unreachable)
  (func (;174;) (type 15)
    i32.const 132042
    i32.const 14
    call 2
    unreachable)
  (func (;175;) (type 15)
    (local i32)
    call 31
    i32.const 0
    call 79
    block  ;; label = @1
      call 147
      local.tee 0
      call 137
      br_if 0 (;@1;)
      local.get 0
      i64.const 10
      call 130
    end
    call 148
    i64.const -1
    call 136
    call 150
    i64.const 1
    call 136
    call 151
    i64.const 1
    call 136)
  (func (;176;) (type 15)
    (local i32 i64 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i64 i32 i64 i64 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 272
    i32.sub
    local.tee 0
    global.set 0
    call 31
    call 67
    call 81
    i32.const 1
    call 80
    call 77
    local.set 1
    local.get 0
    i32.const 1
    i32.store offset=200
    local.get 0
    i32.const 200
    i32.add
    call 71
    local.set 2
    local.get 0
    i32.load offset=200
    call 78
    call 51
    local.set 3
    call 51
    local.set 4
    local.get 0
    call 51
    local.tee 5
    i32.store offset=100
    call 63
    call 66
    local.set 6
    call 166
    call 84
    local.set 7
    local.get 2
    call 15
    local.set 8
    local.get 0
    i32.const 0
    i32.store offset=160
    local.get 0
    local.get 8
    i32.const 2
    i32.shr_u
    local.tee 8
    i32.store offset=156
    local.get 0
    local.get 2
    i32.store offset=152
    local.get 0
    i32.const 168
    i32.add
    i32.const 20
    i32.add
    local.set 9
    local.get 0
    i32.const 168
    i32.add
    i32.const 16
    i32.add
    local.set 10
    local.get 0
    i32.const 200
    i32.add
    i32.const 24
    i32.add
    local.set 11
    i32.const 0
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          loop  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 2
                  local.get 8
                  i32.ge_u
                  br_if 0 (;@7;)
                  local.get 0
                  i32.const 152
                  i32.add
                  i32.const 131463
                  i32.const 8
                  call 156
                  local.tee 2
                  call 15
                  local.set 8
                  local.get 0
                  i32.const 0
                  i32.store8 offset=216
                  local.get 0
                  local.get 8
                  i32.store offset=212
                  local.get 0
                  local.get 2
                  i32.store offset=208
                  local.get 0
                  local.get 8
                  i32.store offset=204
                  local.get 0
                  i32.const 0
                  i32.store offset=200
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 20
                  i32.const 131463
                  i32.const 8
                  call 106
                  local.set 2
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 32
                  i32.const 131463
                  i32.const 8
                  call 106
                  local.set 12
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 131463
                  i32.const 8
                  call 56
                  local.set 13
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 131463
                  i32.const 8
                  call 55
                  local.set 14
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 131463
                  i32.const 8
                  call 160
                  local.set 15
                  local.get 0
                  i32.const 0
                  i32.store8 offset=168
                  local.get 0
                  i32.const 200
                  i32.add
                  local.get 0
                  i32.const 168
                  i32.add
                  i32.const 1
                  i32.const 131463
                  i32.const 8
                  call 105
                  i32.const 2147483646
                  local.set 8
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 0
                        i32.load8_u offset=168
                        br_table 2 (;@8;) 1 (;@9;) 0 (;@10;)
                      end
                      i32.const 131463
                      i32.const 8
                      i32.const 131661
                      i32.const 13
                      call 69
                      unreachable
                    end
                    local.get 0
                    i32.const 200
                    i32.add
                    i32.const 131463
                    i32.const 8
                    call 56
                    local.set 8
                  end
                  local.get 0
                  i32.load offset=204
                  local.get 0
                  i32.load offset=200
                  i32.ne
                  br_if 1 (;@6;)
                  block  ;; label = @8
                    local.get 0
                    i32.load8_u offset=216
                    i32.eqz
                    br_if 0 (;@8;)
                    i32.const 0
                    i32.const 0
                    i32.store offset=142072
                    i32.const 0
                    i32.const 0
                    i32.store8 offset=142076
                  end
                  local.get 0
                  local.get 8
                  i32.store offset=192
                  local.get 0
                  local.get 14
                  i32.store offset=188
                  local.get 0
                  local.get 13
                  i32.store offset=184
                  local.get 0
                  local.get 12
                  i32.store offset=180
                  local.get 0
                  local.get 2
                  i32.store offset=176
                  local.get 0
                  local.get 15
                  i64.store offset=168
                  local.get 7
                  call 49
                  local.set 8
                  local.get 0
                  i32.const 24
                  i32.add
                  i32.const 131471
                  i32.const 9
                  call 88
                  local.get 0
                  i32.const 16
                  i32.add
                  local.get 0
                  i32.load offset=24
                  local.get 0
                  i32.load offset=28
                  local.get 10
                  call 99
                  local.get 0
                  i32.const 8
                  i32.add
                  local.get 0
                  i32.load offset=16
                  local.get 0
                  i32.load offset=20
                  local.get 9
                  call 100
                  local.get 0
                  i32.load offset=12
                  local.set 2
                  local.get 0
                  i32.load offset=8
                  local.set 12
                  call 172
                  local.get 8
                  call 122
                  local.get 12
                  local.get 2
                  call 62
                  local.set 8
                  call 32
                  local.get 0
                  local.get 8
                  i32.store offset=240
                  local.get 0
                  i32.const 240
                  i32.add
                  call 50
                  local.tee 8
                  call 15
                  local.set 2
                  local.get 0
                  i32.const 0
                  i32.store offset=208
                  local.get 0
                  local.get 8
                  i32.store offset=200
                  local.get 0
                  local.get 2
                  i32.const 2
                  i32.shr_u
                  i32.store offset=204
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 131251
                  i32.const 11
                  call 156
                  i32.const 131251
                  i32.const 11
                  call 161
                  local.tee 15
                  i64.const 1
                  i64.gt_u
                  br_if 5 (;@2;)
                  local.get 15
                  i32.wrap_i64
                  br_table 4 (;@3;) 2 (;@5;) 4 (;@3;)
                end
                call 165
                call 137
                i32.eqz
                br_if 5 (;@1;)
                call 165
                call 84
                local.set 8
                local.get 0
                i32.const 88
                i32.add
                i32.const 131651
                i32.const 10
                call 88
                local.get 0
                local.get 3
                i32.store offset=252
                local.get 0
                local.get 8
                i32.store offset=248
                local.get 0
                local.get 0
                i32.load offset=92
                local.tee 2
                i32.store offset=244
                local.get 0
                local.get 0
                i32.load offset=88
                local.tee 12
                i32.store offset=240
                local.get 0
                call 172
                local.tee 15
                i64.store offset=152
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 3
                        call 89
                        br_table 1 (;@9;) 2 (;@8;) 0 (;@10;)
                      end
                      local.get 0
                      i32.const 64
                      i32.add
                      local.get 12
                      local.get 2
                      local.get 8
                      local.get 3
                      call 87
                      local.get 0
                      i32.load offset=68
                      local.set 8
                      local.get 0
                      i32.load offset=64
                      local.set 2
                      local.get 15
                      call 63
                      call 122
                      local.get 2
                      local.get 8
                      call 103
                      local.set 8
                      br 2 (;@7;)
                    end
                    local.get 15
                    local.get 8
                    call 122
                    local.get 12
                    local.get 2
                    call 103
                    local.set 8
                    br 1 (;@7;)
                  end
                  local.get 0
                  i32.const 168
                  i32.add
                  local.get 3
                  call 115
                  local.get 0
                  i64.load offset=168
                  local.set 15
                  local.get 0
                  local.get 0
                  i32.const 180
                  i32.add
                  local.tee 14
                  i32.store offset=220
                  local.get 0
                  local.get 0
                  i32.const 168
                  i32.add
                  i32.const 8
                  i32.add
                  local.tee 3
                  i32.store offset=216
                  local.get 0
                  local.get 15
                  i64.store offset=208
                  local.get 0
                  local.get 0
                  i32.const 256
                  i32.add
                  local.tee 13
                  i32.store offset=232
                  local.get 0
                  local.get 13
                  i32.store offset=228
                  local.get 0
                  local.get 2
                  i32.store offset=204
                  local.get 0
                  local.get 12
                  i32.store offset=200
                  local.get 0
                  local.get 0
                  i32.const 152
                  i32.add
                  i32.store offset=224
                  block  ;; label = @8
                    local.get 15
                    i64.const 0
                    i64.ne
                    br_if 0 (;@8;)
                    local.get 0
                    i32.const 72
                    i32.add
                    local.get 12
                    local.get 2
                    local.get 3
                    local.get 14
                    call 101
                    local.get 0
                    i32.load offset=76
                    local.set 2
                    local.get 0
                    i32.load offset=72
                    local.set 12
                    call 122
                    local.set 13
                    local.get 0
                    i64.load offset=152
                    local.get 8
                    local.get 13
                    local.get 12
                    local.get 2
                    call 103
                    local.set 8
                    br 1 (;@7;)
                  end
                  local.get 0
                  i32.const 80
                  i32.add
                  local.get 12
                  local.get 2
                  local.get 8
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 8
                  i32.add
                  call 98
                  local.get 0
                  i32.load offset=84
                  local.set 8
                  local.get 0
                  i32.load offset=80
                  local.set 2
                  call 63
                  local.set 12
                  call 122
                  local.set 13
                  local.get 0
                  i64.load offset=152
                  local.get 12
                  local.get 13
                  local.get 2
                  local.get 8
                  call 103
                  local.set 8
                end
                call 32
                local.get 0
                local.get 8
                i32.store offset=148
                i32.const 0
                local.set 2
                local.get 0
                i32.const 148
                i32.add
                call 50
                local.tee 8
                call 15
                local.set 12
                local.get 0
                i32.const 0
                i32.store offset=160
                local.get 0
                local.get 8
                i32.store offset=152
                local.get 0
                local.get 12
                i32.const 2
                i32.shr_u
                i32.store offset=156
                local.get 0
                i32.const 152
                i32.add
                i32.const 131251
                i32.const 11
                call 156
                call 49
                local.set 12
                call 51
                local.set 3
                local.get 12
                call 15
                local.set 8
                local.get 0
                i32.const 0
                i32.store8 offset=216
                local.get 0
                local.get 8
                i32.store offset=212
                local.get 0
                local.get 12
                i32.store offset=208
                local.get 0
                local.get 8
                i32.store offset=204
                local.get 0
                i32.const 0
                i32.store offset=200
                loop  ;; label = @7
                  block  ;; label = @8
                    local.get 8
                    local.get 2
                    i32.ne
                    br_if 0 (;@8;)
                    local.get 0
                    i32.load8_u offset=216
                    i32.eqz
                    br_if 7 (;@1;)
                    i32.const 0
                    i32.const 0
                    i32.store offset=142072
                    i32.const 0
                    i32.const 0
                    i32.store8 offset=142076
                    br 7 (;@1;)
                  end
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 131251
                  i32.const 11
                  call 56
                  local.set 8
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 131251
                  i32.const 11
                  call 160
                  local.set 15
                  local.get 0
                  local.get 0
                  i32.const 200
                  i32.add
                  i32.const 131251
                  i32.const 11
                  call 55
                  i32.store offset=180
                  local.get 0
                  local.get 8
                  i32.store offset=176
                  local.get 0
                  local.get 15
                  i64.store offset=168
                  local.get 3
                  local.get 0
                  i32.const 168
                  i32.add
                  call 117
                  local.get 0
                  i32.load offset=200
                  local.set 2
                  local.get 0
                  i32.load offset=204
                  local.set 8
                  br 0 (;@7;)
                end
              end
              i32.const 131463
              i32.const 8
              i32.const 131072
              i32.const 14
              call 69
              unreachable
            end
            local.get 0
            i32.const 200
            i32.add
            local.get 0
            i32.const 168
            i32.add
            call 169
            i32.const 131518
            i32.const 31
            call 70
            local.set 8
            local.get 0
            i32.load offset=216
            local.tee 2
            local.get 8
            call 60
            call 165
            call 84
            local.tee 12
            local.get 8
            i32.const -25
            call 33
            block  ;; label = @5
              i32.const -25
              call 15
              i32.eqz
              br_if 0 (;@5;)
              local.get 12
              local.get 8
              call 42
              local.tee 2
              call 33
            end
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 0
                      i32.load offset=180
                      call 15
                      i32.const 32
                      i32.ne
                      br_if 0 (;@9;)
                      local.get 11
                      i64.const 0
                      i64.store
                      local.get 0
                      i32.const 200
                      i32.add
                      i32.const 16
                      i32.add
                      i64.const 0
                      i64.store
                      local.get 0
                      i32.const 200
                      i32.add
                      i32.const 8
                      i32.add
                      i64.const 0
                      i64.store
                      local.get 0
                      i64.const 0
                      i64.store offset=200
                      local.get 0
                      i32.load offset=180
                      i32.const 0
                      local.get 0
                      i32.const 200
                      i32.add
                      i32.const 32
                      call 113
                      drop
                      local.get 0
                      i32.const 200
                      i32.add
                      i32.const 131086
                      i32.const 32
                      call 202
                      i32.eqz
                      br_if 1 (;@8;)
                    end
                    block  ;; label = @9
                      local.get 0
                      i32.load offset=184
                      call 168
                      call 140
                      local.tee 8
                      call 34
                      i32.const 1
                      i32.lt_s
                      br_if 0 (;@9;)
                      local.get 0
                      i32.load offset=188
                      local.get 8
                      call 35
                      i32.const 0
                      i32.gt_s
                      br_if 2 (;@7;)
                    end
                    block  ;; label = @9
                      local.get 6
                      local.get 0
                      i32.load offset=180
                      local.get 2
                      call 159
                      i32.eqz
                      br_if 0 (;@9;)
                      local.get 0
                      i64.load offset=168
                      local.set 15
                      i32.const 131933
                      i32.const 38
                      call 143
                      local.tee 8
                      local.get 1
                      call 142
                      local.get 8
                      local.get 15
                      call 142
                      br 3 (;@6;)
                    end
                    local.get 0
                    i32.load offset=176
                    call 49
                    local.set 2
                    local.get 0
                    i32.load offset=180
                    call 49
                    local.set 12
                    local.get 0
                    i32.load offset=184
                    call 49
                    local.set 13
                    local.get 0
                    i32.load offset=188
                    call 41
                    local.set 14
                    local.get 0
                    i64.load offset=168
                    local.set 15
                    i32.const 131807
                    i32.const 22
                    call 143
                    local.tee 8
                    local.get 1
                    call 142
                    call 51
                    call 49
                    local.tee 16
                    local.get 2
                    call 4
                    drop
                    local.get 8
                    local.get 16
                    call 54
                    local.get 8
                    local.get 12
                    call 95
                    local.get 8
                    local.get 13
                    call 95
                    local.get 8
                    local.get 14
                    call 97
                    local.get 8
                    local.get 15
                    call 142
                    local.get 8
                    call 51
                    call 36
                    local.get 0
                    i32.const 200
                    i32.add
                    local.get 0
                    i32.const 168
                    i32.add
                    call 169
                    local.get 0
                    local.get 0
                    i32.load offset=208
                    local.tee 8
                    i32.const 24
                    i32.shl
                    local.get 8
                    i32.const 65280
                    i32.and
                    i32.const 8
                    i32.shl
                    i32.or
                    local.get 8
                    i32.const 8
                    i32.shr_u
                    i32.const 65280
                    i32.and
                    local.get 8
                    i32.const 24
                    i32.shr_u
                    i32.or
                    i32.or
                    i32.store offset=240
                    local.get 0
                    local.get 0
                    i32.load offset=212
                    local.tee 8
                    i32.const 24
                    i32.shl
                    local.get 8
                    i32.const 65280
                    i32.and
                    i32.const 8
                    i32.shl
                    i32.or
                    local.get 8
                    i32.const 8
                    i32.shr_u
                    i32.const 65280
                    i32.and
                    local.get 8
                    i32.const 24
                    i32.shr_u
                    i32.or
                    i32.or
                    i32.store offset=244
                    local.get 0
                    local.get 0
                    i32.load offset=216
                    local.tee 8
                    i32.const 24
                    i32.shl
                    local.get 8
                    i32.const 65280
                    i32.and
                    i32.const 8
                    i32.shl
                    i32.or
                    local.get 8
                    i32.const 8
                    i32.shr_u
                    i32.const 65280
                    i32.and
                    local.get 8
                    i32.const 24
                    i32.shr_u
                    i32.or
                    i32.or
                    i32.store offset=248
                    local.get 0
                    local.get 0
                    i32.load offset=220
                    local.tee 8
                    i32.const 24
                    i32.shl
                    local.get 8
                    i32.const 65280
                    i32.and
                    i32.const 8
                    i32.shl
                    i32.or
                    local.get 8
                    i32.const 8
                    i32.shr_u
                    i32.const 65280
                    i32.and
                    local.get 8
                    i32.const 24
                    i32.shr_u
                    i32.or
                    i32.or
                    i32.store offset=252
                    local.get 0
                    local.get 0
                    i32.load offset=224
                    local.tee 8
                    i32.const 24
                    i32.shl
                    local.get 8
                    i32.const 65280
                    i32.and
                    i32.const 8
                    i32.shl
                    i32.or
                    local.get 8
                    i32.const 8
                    i32.shr_u
                    i32.const 65280
                    i32.and
                    local.get 8
                    i32.const 24
                    i32.shr_u
                    i32.or
                    i32.or
                    i32.store offset=264
                    local.get 0
                    local.get 0
                    i64.load offset=200
                    local.tee 15
                    i64.const 56
                    i64.shl
                    local.get 15
                    i64.const 65280
                    i64.and
                    i64.const 40
                    i64.shl
                    i64.or
                    local.get 15
                    i64.const 16711680
                    i64.and
                    i64.const 24
                    i64.shl
                    local.get 15
                    i64.const 4278190080
                    i64.and
                    i64.const 8
                    i64.shl
                    i64.or
                    i64.or
                    local.get 15
                    i64.const 8
                    i64.shr_u
                    i64.const 4278190080
                    i64.and
                    local.get 15
                    i64.const 24
                    i64.shr_u
                    i64.const 16711680
                    i64.and
                    i64.or
                    local.get 15
                    i64.const 40
                    i64.shr_u
                    i64.const 65280
                    i64.and
                    local.get 15
                    i64.const 56
                    i64.shr_u
                    i64.or
                    i64.or
                    i64.or
                    i64.store offset=256 align=4
                    local.get 4
                    local.get 0
                    i32.const 240
                    i32.add
                    i32.const 28
                    call 13
                    drop
                    local.get 0
                    local.get 0
                    i64.load offset=184
                    i64.store offset=208
                    local.get 0
                    i64.const 0
                    i64.store offset=200
                    local.get 3
                    local.get 0
                    i32.const 200
                    i32.add
                    call 117
                    br 3 (;@5;)
                  end
                  local.get 0
                  i64.load offset=168
                  local.set 15
                  i32.const 131901
                  i32.const 32
                  call 143
                  local.tee 8
                  local.get 1
                  call 142
                  local.get 8
                  local.get 15
                  call 142
                  br 1 (;@6;)
                end
                local.get 0
                i64.load offset=168
                local.set 15
                i32.const 131786
                i32.const 21
                call 143
                local.tee 8
                local.get 1
                call 142
                local.get 8
                local.get 15
                call 142
              end
              local.get 8
              call 51
              call 36
              local.get 0
              i64.load offset=168
              local.set 15
              local.get 0
              i64.load offset=184
              local.set 17
              local.get 0
              i32.load offset=180
              local.set 8
              local.get 0
              i32.load offset=176
              local.set 2
              call 27
              local.set 18
              local.get 2
              call 49
              local.set 2
              local.get 8
              call 49
              local.set 8
              local.get 0
              i32.const 1
              i32.store8 offset=136
              local.get 0
              local.get 17
              i64.store offset=128
              local.get 0
              local.get 8
              i32.store offset=124
              local.get 0
              local.get 2
              i32.store offset=120
              local.get 0
              local.get 15
              i64.store offset=112
              local.get 0
              local.get 18
              i64.store offset=104
              local.get 5
              local.get 0
              i32.const 104
              i32.add
              call 116
            end
            local.get 0
            i32.load offset=156
            local.set 8
            local.get 0
            i32.load offset=160
            local.set 2
            br 0 (;@4;)
          end
        end
        i32.const 131549
        i32.const 23
        call 61
        unreachable
      end
      i32.const 131251
      i32.const 11
      i32.const 131674
      i32.const 18
      call 69
      unreachable
    end
    local.get 0
    i32.const 200
    i32.add
    i32.const 8
    i32.add
    local.set 2
    local.get 0
    i32.const 168
    i32.add
    i32.const 28
    i32.add
    local.set 19
    local.get 0
    i32.const 168
    i32.add
    i32.const 8
    i32.add
    local.set 20
    local.get 0
    i32.const 252
    i32.add
    local.set 21
    local.get 0
    i32.const 240
    i32.add
    i32.const 8
    i32.add
    local.set 22
    call 167
    call 84
    local.set 23
    local.get 4
    call 15
    local.set 16
    local.get 3
    call 15
    local.set 24
    local.get 0
    i32.const 200
    i32.add
    i32.const 24
    i32.add
    local.set 10
    i32.const 0
    local.set 8
    i32.const 0
    local.set 12
    block  ;; label = @1
      loop  ;; label = @2
        local.get 8
        i32.const 28
        i32.add
        local.tee 5
        local.get 16
        i32.gt_u
        br_if 1 (;@1;)
        local.get 10
        i32.const 0
        i32.store
        local.get 0
        i32.const 200
        i32.add
        i32.const 16
        i32.add
        local.tee 13
        i64.const 0
        i64.store
        local.get 2
        i64.const 0
        i64.store
        local.get 0
        i64.const 0
        i64.store offset=200
        local.get 4
        local.get 8
        local.get 0
        i32.const 200
        i32.add
        i32.const 28
        call 113
        drop
        local.get 12
        i32.const 16
        i32.add
        local.tee 11
        local.get 24
        i32.gt_u
        br_if 1 (;@1;)
        local.get 10
        i32.load
        local.set 9
        local.get 13
        i64.load
        local.set 15
        local.get 2
        i32.load
        local.set 13
        local.get 0
        i32.load offset=212
        local.set 14
        local.get 0
        i32.load offset=204
        local.set 8
        local.get 0
        i32.load offset=200
        local.set 7
        local.get 2
        i64.const 0
        i64.store
        local.get 0
        i64.const 0
        i64.store offset=200
        local.get 3
        local.get 12
        local.get 0
        i32.const 200
        i32.add
        i32.const 16
        call 113
        drop
        local.get 0
        i32.const 0
        i32.store offset=168
        local.get 0
        i32.const 200
        i32.add
        local.get 0
        i32.const 168
        i32.add
        call 93
        local.set 12
        local.get 0
        i32.const 200
        i32.add
        local.get 0
        i32.const 168
        i32.add
        call 94
        local.set 17
        local.get 0
        local.get 0
        i32.const 200
        i32.add
        local.get 0
        i32.const 168
        i32.add
        call 93
        i32.store offset=252
        local.get 0
        local.get 12
        i32.store offset=248
        local.get 0
        local.get 17
        i64.store offset=240
        block  ;; label = @3
          block  ;; label = @4
            local.get 8
            i32.const 24
            i32.shl
            local.get 8
            i32.const 65280
            i32.and
            i32.const 8
            i32.shl
            i32.or
            local.get 8
            i32.const 8
            i32.shr_u
            i32.const 65280
            i32.and
            local.get 8
            i32.const 24
            i32.shr_u
            i32.or
            i32.or
            local.tee 12
            call 64
            br_if 0 (;@4;)
            call 51
            local.set 13
            call 51
            local.set 14
            call 51
            local.set 8
            local.get 0
            i32.load offset=248
            call 49
            local.set 7
            local.get 0
            local.get 0
            i32.load offset=252
            call 41
            i32.store offset=212
            local.get 0
            i64.const 0
            i64.store offset=200
            local.get 0
            local.get 7
            i32.store offset=208
            local.get 8
            local.get 0
            i32.const 200
            i32.add
            call 117
            local.get 12
            local.get 8
            i64.const 0
            local.get 13
            local.get 14
            call 37
            drop
            br 1 (;@3;)
          end
          local.get 23
          call 49
          local.set 25
          local.get 0
          i32.const 56
          i32.add
          i32.const 131644
          i32.const 7
          call 88
          local.get 0
          i32.load offset=60
          local.set 6
          local.get 0
          i32.load offset=56
          local.set 26
          call 51
          call 49
          local.tee 8
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
          call 4
          drop
          local.get 8
          local.get 12
          call 4
          drop
          local.get 13
          i32.const 24
          i32.shl
          local.get 13
          i32.const 65280
          i32.and
          i32.const 8
          i32.shl
          i32.or
          local.get 13
          i32.const 8
          i32.shr_u
          i32.const 65280
          i32.and
          local.get 13
          i32.const 24
          i32.shr_u
          i32.or
          i32.or
          local.get 8
          call 60
          local.get 14
          i32.const 24
          i32.shl
          local.get 14
          i32.const 65280
          i32.and
          i32.const 8
          i32.shl
          i32.or
          local.get 14
          i32.const 8
          i32.shr_u
          i32.const 65280
          i32.and
          local.get 14
          i32.const 24
          i32.shr_u
          i32.or
          i32.or
          local.get 8
          call 58
          local.get 15
          i64.const 56
          i64.shl
          local.get 15
          i64.const 65280
          i64.and
          i64.const 40
          i64.shl
          i64.or
          local.get 15
          i64.const 16711680
          i64.and
          i64.const 24
          i64.shl
          local.get 15
          i64.const 4278190080
          i64.and
          i64.const 8
          i64.shl
          i64.or
          i64.or
          local.get 15
          i64.const 8
          i64.shr_u
          i64.const 4278190080
          i64.and
          local.get 15
          i64.const 24
          i64.shr_u
          i64.const 16711680
          i64.and
          i64.or
          local.get 15
          i64.const 40
          i64.shr_u
          i64.const 65280
          i64.and
          local.get 15
          i64.const 56
          i64.shr_u
          i64.or
          i64.or
          i64.or
          local.get 8
          call 125
          block  ;; label = @4
            block  ;; label = @5
              local.get 9
              i32.const -16777345
              i32.ne
              br_if 0 (;@5;)
              local.get 8
              i32.const 0
              call 163
              br 1 (;@4;)
            end
            local.get 8
            i32.const 1
            call 163
            local.get 9
            i32.const 24
            i32.shl
            local.get 9
            i32.const 65280
            i32.and
            i32.const 8
            i32.shl
            i32.or
            local.get 9
            i32.const 8
            i32.shr_u
            i32.const 65280
            i32.and
            local.get 9
            i32.const 24
            i32.shr_u
            i32.or
            i32.or
            local.get 8
            call 60
          end
          local.get 6
          local.get 8
          call 54
          local.get 0
          i32.const 48
          i32.add
          local.get 26
          local.get 6
          local.get 1
          call 96
          local.get 0
          local.get 25
          i32.store offset=192
          local.get 0
          local.get 21
          i32.store offset=188
          local.get 0
          local.get 22
          i32.store offset=184
          local.get 0
          local.get 0
          i64.load offset=48
          i64.store offset=168
          local.get 0
          i64.const 0
          i64.store offset=176
          local.get 0
          call 172
          i64.store offset=152
          local.get 2
          i32.const 8
          i32.add
          local.get 20
          i32.const 8
          i32.add
          i64.load
          i64.store
          local.get 2
          local.get 20
          i64.load
          i64.store
          local.get 0
          local.get 19
          i32.store offset=232
          local.get 0
          local.get 19
          i32.store offset=228
          local.get 0
          local.get 0
          i32.load offset=172
          local.tee 8
          i32.store offset=204
          local.get 0
          local.get 0
          i32.load offset=168
          local.tee 12
          i32.store offset=200
          local.get 0
          i64.load offset=208
          local.set 15
          local.get 0
          local.get 0
          i32.const 152
          i32.add
          i32.store offset=224
          local.get 0
          i32.load offset=192
          local.set 13
          block  ;; label = @4
            block  ;; label = @5
              local.get 15
              i64.const 0
              i64.ne
              br_if 0 (;@5;)
              local.get 0
              i32.const 32
              i32.add
              local.get 12
              local.get 8
              local.get 0
              i32.load offset=216
              local.get 0
              i32.load offset=220
              call 101
              local.get 0
              i32.load offset=36
              local.set 8
              local.get 0
              i32.load offset=32
              local.set 12
              call 122
              local.set 14
              local.get 0
              i64.load offset=152
              local.get 13
              local.get 14
              local.get 12
              local.get 8
              call 104
              br 1 (;@4;)
            end
            local.get 0
            i32.const 40
            i32.add
            local.get 12
            local.get 8
            local.get 13
            local.get 2
            call 98
            local.get 0
            i32.load offset=44
            local.set 8
            local.get 0
            i32.load offset=40
            local.set 12
            call 63
            local.set 13
            call 122
            local.set 14
            local.get 0
            i64.load offset=152
            local.get 13
            local.get 14
            local.get 12
            local.get 8
            call 104
          end
          call 32
        end
        local.get 5
        local.set 8
        local.get 11
        local.set 12
        br 0 (;@2;)
      end
    end
    local.get 0
    i32.const 100
    i32.add
    call 149
    local.get 0
    i32.const 272
    i32.add
    global.set 0)
  (func (;177;) (type 15)
    (local i32 i32 i32 i64 i32 i64 i64 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 240
    i32.sub
    local.tee 0
    global.set 0
    call 31
    call 67
    i32.const 0
    local.set 1
    i32.const 0
    call 79
    local.get 0
    i32.const 40
    i32.add
    call 154
    block  ;; label = @1
      local.get 0
      i64.load offset=40
      i64.const 0
      i64.ne
      br_if 0 (;@1;)
      local.get 0
      i32.load offset=56
      local.set 2
      local.get 0
      i32.const 64
      i32.add
      call 150
      call 86
      call 152
      call 150
      call 86
      local.set 3
      call 151
      local.set 4
      local.get 4
      local.get 3
      i64.const 1
      i64.add
      local.tee 5
      local.get 4
      call 86
      local.tee 6
      local.get 6
      local.get 3
      i64.eq
      select
      call 130
      call 150
      local.get 5
      call 130
      local.get 0
      i32.load offset=68
      local.tee 7
      call 135
      local.set 8
      local.get 0
      i32.load offset=64
      local.set 9
      i32.const 1
      local.set 4
      block  ;; label = @2
        loop  ;; label = @3
          local.get 1
          i32.const 1
          i32.and
          br_if 1 (;@2;)
          local.get 4
          local.get 8
          i32.gt_u
          br_if 1 (;@2;)
          local.get 9
          local.get 4
          call 131
          call 127
          local.get 4
          local.get 8
          i32.ge_u
          local.set 1
          local.get 4
          local.get 4
          local.get 8
          i32.lt_u
          i32.add
          local.set 4
          br 0 (;@3;)
        end
      end
      local.get 7
      i64.const 0
      call 130
      call 51
      local.set 10
      call 51
      local.set 11
      local.get 2
      call 15
      local.set 4
      local.get 0
      i32.const 0
      i32.store offset=192
      local.get 0
      local.get 4
      i32.const 2
      i32.shr_u
      local.tee 4
      i32.store offset=188
      local.get 0
      local.get 2
      i32.store offset=184
      i32.const 0
      local.set 8
      loop  ;; label = @2
        block  ;; label = @3
          local.get 8
          local.get 4
          i32.lt_u
          br_if 0 (;@3;)
          call 166
          call 84
          local.set 12
          local.get 0
          local.get 10
          i32.store offset=176
          local.get 0
          i32.const 32
          i32.add
          i32.const 131480
          i32.const 14
          call 88
          local.get 0
          i32.load offset=36
          local.set 13
          local.get 0
          i32.load offset=32
          local.set 14
          call 51
          call 49
          local.set 4
          local.get 0
          local.get 10
          call 15
          i32.store offset=192
          local.get 0
          i32.const 0
          i32.store offset=188
          local.get 0
          local.get 0
          i32.const 176
          i32.add
          i32.store offset=184
          block  ;; label = @4
            loop  ;; label = @5
              local.get 0
              i32.const 200
              i32.add
              local.get 0
              i32.const 184
              i32.add
              call 145
              local.get 0
              i32.load8_u offset=232
              local.tee 8
              i32.const 2
              i32.eq
              br_if 1 (;@4;)
              local.get 0
              i32.load offset=228
              local.set 1
              local.get 0
              i32.load offset=224
              local.set 9
              local.get 0
              i32.load offset=220
              local.set 2
              local.get 0
              i32.load offset=216
              local.set 7
              local.get 0
              i64.load offset=208
              local.set 3
              local.get 0
              i64.load offset=200
              local.get 4
              call 125
              local.get 3
              local.get 4
              call 125
              local.get 7
              local.get 4
              call 60
              local.get 2
              local.get 4
              call 60
              local.get 9
              local.get 4
              call 60
              local.get 1
              local.get 4
              call 58
              local.get 4
              local.get 8
              call 163
              br 0 (;@5;)
            end
          end
          local.get 13
          local.get 4
          call 54
          local.get 0
          local.get 11
          i32.store offset=172
          local.get 0
          local.get 12
          i32.store offset=168
          local.get 0
          local.get 13
          i32.store offset=164
          local.get 0
          local.get 14
          i32.store offset=160
          local.get 0
          call 172
          local.tee 3
          i64.store offset=176
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 11
                  call 89
                  br_table 1 (;@6;) 2 (;@5;) 0 (;@7;)
                end
                local.get 0
                i32.const 8
                i32.add
                local.get 14
                local.get 13
                local.get 12
                local.get 11
                call 87
                local.get 0
                i32.load offset=12
                local.set 4
                local.get 0
                i32.load offset=8
                local.set 8
                local.get 3
                call 63
                call 122
                local.get 8
                local.get 4
                call 104
                br 2 (;@4;)
              end
              local.get 3
              local.get 12
              call 122
              local.get 14
              local.get 13
              call 104
              br 1 (;@4;)
            end
            local.get 0
            i32.const 184
            i32.add
            local.get 11
            call 115
            local.get 0
            i64.load offset=184
            local.set 3
            local.get 0
            local.get 0
            i32.const 196
            i32.add
            local.tee 8
            i32.store offset=220
            local.get 0
            local.get 0
            i32.const 184
            i32.add
            i32.const 8
            i32.add
            local.tee 1
            i32.store offset=216
            local.get 0
            local.get 3
            i64.store offset=208
            local.get 0
            local.get 0
            i32.const 176
            i32.add
            local.tee 4
            i32.store offset=232
            local.get 0
            local.get 4
            i32.store offset=228
            local.get 0
            local.get 13
            i32.store offset=204
            local.get 0
            local.get 14
            i32.store offset=200
            local.get 0
            local.get 0
            i32.const 176
            i32.add
            i32.store offset=224
            block  ;; label = @5
              local.get 3
              i64.const 0
              i64.ne
              br_if 0 (;@5;)
              local.get 0
              i32.const 16
              i32.add
              local.get 14
              local.get 13
              local.get 1
              local.get 8
              call 101
              local.get 0
              i32.load offset=20
              local.set 4
              local.get 0
              i32.load offset=16
              local.set 8
              call 122
              local.set 1
              local.get 0
              i64.load offset=176
              local.get 12
              local.get 1
              local.get 8
              local.get 4
              call 104
              br 1 (;@4;)
            end
            local.get 0
            i32.const 24
            i32.add
            local.get 14
            local.get 13
            local.get 12
            local.get 0
            i32.const 200
            i32.add
            i32.const 8
            i32.add
            call 98
            local.get 0
            i32.load offset=28
            local.set 4
            local.get 0
            i32.load offset=24
            local.set 8
            call 63
            local.set 1
            call 122
            local.set 9
            local.get 0
            i64.load offset=176
            local.get 1
            local.get 9
            local.get 8
            local.get 4
            call 104
          end
          call 32
          br 2 (;@1;)
        end
        local.get 0
        i32.const 184
        i32.add
        call 171
        local.set 5
        local.get 0
        i32.const 184
        i32.add
        call 171
        local.set 3
        local.get 0
        i32.const 184
        i32.add
        call 170
        local.set 4
        local.get 0
        i32.const 184
        i32.add
        call 170
        local.set 8
        local.get 0
        i32.const 184
        i32.add
        i32.const 131463
        i32.const 8
        call 156
        call 49
        local.set 1
        local.get 0
        i32.const 184
        i32.add
        i32.const 131463
        i32.const 8
        call 156
        call 57
        local.set 9
        local.get 4
        call 49
        drop
        local.get 8
        call 49
        drop
        local.get 1
        call 49
        local.set 2
        local.get 9
        call 41
        local.set 13
        call 166
        call 84
        local.set 7
        block  ;; label = @3
          block  ;; label = @4
            call 63
            call 66
            local.get 7
            local.get 2
            call 159
            i32.eqz
            br_if 0 (;@4;)
            local.get 3
            call 164
            local.set 2
            local.get 0
            i32.const 0
            i32.store8 offset=232
            local.get 0
            local.get 9
            i32.store offset=228
            local.get 0
            local.get 1
            i32.store offset=224
            local.get 0
            local.get 8
            i32.store offset=220
            local.get 0
            local.get 4
            i32.store offset=216
            local.get 0
            local.get 3
            i64.store offset=208
            local.get 0
            local.get 5
            i64.store offset=200
            local.get 2
            local.get 0
            i32.const 200
            i32.add
            call 124
            i32.const 131766
            i32.const 20
            call 143
            local.tee 4
            local.get 3
            call 142
            local.get 4
            call 51
            call 36
            br 1 (;@3;)
          end
          local.get 0
          i32.const 0
          i32.store8 offset=232
          local.get 0
          local.get 9
          i32.store offset=228
          local.get 0
          local.get 1
          i32.store offset=224
          local.get 0
          local.get 8
          i32.store offset=220
          local.get 0
          local.get 4
          i32.store offset=216
          local.get 0
          local.get 3
          i64.store offset=208
          local.get 0
          local.get 5
          i64.store offset=200
          local.get 10
          local.get 0
          i32.const 200
          i32.add
          call 116
          local.get 0
          local.get 13
          i32.store offset=212
          local.get 0
          i64.const 0
          i64.store offset=200
          local.get 0
          local.get 2
          i32.store offset=208
          local.get 11
          local.get 0
          i32.const 200
          i32.add
          call 117
        end
        local.get 0
        i32.load offset=188
        local.set 4
        local.get 0
        i32.load offset=192
        local.set 8
        br 0 (;@2;)
      end
    end
    local.get 0
    i32.const 240
    i32.add
    global.set 0)
  (func (;178;) (type 15)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 31
    call 67
    call 81
    i32.const 0
    call 80
    local.get 0
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 0
    i32.const 12
    i32.add
    call 73
    local.get 0
    i32.load offset=4
    local.set 1
    local.get 0
    i32.load
    local.set 2
    local.get 0
    i32.load offset=12
    call 78
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          br_if 0 (;@3;)
          local.get 1
          call 64
          i32.eqz
          br_if 1 (;@2;)
          call 165
          local.get 1
          call 139
          br 2 (;@1;)
        end
        call 165
        call 127
        br 1 (;@1;)
      end
      i32.const 131572
      i32.const 35
      call 61
      unreachable
    end
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;179;) (type 15)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 31
    call 67
    call 81
    i32.const 0
    call 80
    local.get 0
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 0
    i32.const 12
    i32.add
    call 73
    local.get 0
    i32.load offset=4
    local.set 1
    local.get 0
    i32.load
    local.set 2
    local.get 0
    i32.load offset=12
    call 78
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          br_if 0 (;@3;)
          local.get 1
          call 64
          i32.eqz
          br_if 1 (;@2;)
          call 167
          local.get 1
          call 139
          br 2 (;@1;)
        end
        call 167
        call 127
        br 1 (;@1;)
      end
      i32.const 131607
      i32.const 37
      call 61
      unreachable
    end
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;180;) (type 15)
    (local i32 i64 i32)
    global.get 0
    i32.const 48
    i32.sub
    local.tee 0
    global.set 0
    call 31
    call 67
    i32.const 1
    call 79
    local.get 0
    call 77
    local.tee 1
    call 164
    call 132
    local.get 0
    call 51
    local.tee 2
    i32.store offset=44
    local.get 2
    local.get 0
    call 116
    local.get 0
    i32.const 44
    i32.add
    call 149
    local.get 1
    call 164
    call 127
    local.get 0
    i32.const 48
    i32.add
    global.set 0)
  (func (;181;) (type 15)
    (local i32 i32 i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 0
    global.set 0
    call 31
    call 67
    call 81
    i32.const 0
    call 80
    local.get 0
    i32.const 0
    i32.store offset=12
    local.get 0
    local.get 0
    i32.const 12
    i32.add
    call 73
    local.get 0
    i32.load offset=4
    local.set 1
    local.get 0
    i32.load
    local.set 2
    local.get 0
    i32.load offset=12
    call 78
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        br_if 0 (;@2;)
        call 166
        local.get 1
        call 139
        br 1 (;@1;)
      end
      call 166
      call 127
    end
    local.get 0
    i32.const 16
    i32.add
    global.set 0)
  (func (;182;) (type 15)
    call 31
    i32.const 0
    call 79
    call 165
    call 83)
  (func (;183;) (type 15)
    call 31
    i32.const 0
    call 79
    call 167
    call 83)
  (func (;184;) (type 15)
    call 31
    i32.const 0
    call 79
    call 166
    call 83)
  (func (;185;) (type 15)
    (local i32)
    call 31
    call 67
    i32.const 1
    call 79
    block  ;; label = @1
      call 76
      local.tee 0
      br_if 0 (;@1;)
      i32.const 131358
      i32.const 37
      call 61
      unreachable
    end
    call 147
    local.get 0
    i64.extend_i32_u
    call 130)
  (func (;186;) (type 15)
    (local i64)
    call 31
    call 67
    i32.const 1
    call 79
    block  ;; label = @1
      call 77
      local.tee 0
      i64.const 0
      i64.ne
      br_if 0 (;@1;)
      i32.const 131395
      i32.const 47
      call 61
      unreachable
    end
    call 148
    local.get 0
    call 130)
  (func (;187;) (type 15)
    (local i32 i64 i32 i64 i64 i64)
    global.get 0
    i32.const 176
    i32.sub
    local.tee 0
    global.set 0
    call 31
    i32.const 0
    call 79
    local.get 0
    call 150
    call 86
    local.tee 1
    call 152
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          local.get 1
          local.get 1
          call 146
          i32.eqz
          br_if 0 (;@3;)
          local.get 0
          i32.load offset=88
          i32.eqz
          br_if 0 (;@3;)
          local.get 0
          call 48
          local.set 2
          call 27
          local.tee 3
          local.get 2
          i64.load
          local.tee 4
          i64.lt_u
          br_if 0 (;@3;)
          local.get 3
          local.get 4
          i64.sub
          i64.const 10
          i64.gt_u
          br_if 1 (;@2;)
        end
        local.get 0
        i64.const 1
        i64.store offset=144
        br 1 (;@1;)
      end
      call 51
      local.set 2
      local.get 0
      i32.const 1
      i32.store offset=100
      local.get 0
      local.get 0
      i32.store offset=96
      loop  ;; label = @2
        local.get 0
        i32.const 104
        i32.add
        local.get 0
        i32.const 96
        i32.add
        call 44
        block  ;; label = @3
          local.get 0
          i32.load8_u offset=136
          i32.const 2
          i32.ne
          br_if 0 (;@3;)
          local.get 0
          local.get 2
          i32.store offset=160
          local.get 0
          local.get 1
          i64.store offset=152
          local.get 0
          i64.const 0
          i64.store offset=144
          br 2 (;@1;)
        end
        local.get 0
        i64.load offset=104
        local.set 3
        local.get 0
        i64.load offset=112
        local.set 4
        local.get 0
        i64.load offset=120
        local.set 5
        local.get 0
        local.get 0
        i64.load offset=128
        i64.store offset=168
        local.get 0
        local.get 5
        i64.store offset=160
        local.get 0
        local.get 4
        i64.store offset=152
        local.get 0
        local.get 3
        i64.store offset=144
        local.get 2
        local.get 0
        i32.const 144
        i32.add
        call 112
        br 0 (;@2;)
      end
    end
    local.get 0
    i32.const 144
    i32.add
    call 82
    local.get 0
    i32.const 176
    i32.add
    global.set 0)
  (func (;188;) (type 15)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    call 31
    i32.const 0
    call 79
    local.get 0
    i32.const 8
    i32.add
    call 154
    local.get 0
    i32.const 8
    i32.add
    call 82
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;189;) (type 15)
    (local i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    call 31
    i32.const 1
    call 79
    local.get 0
    i32.const 8
    i32.add
    call 77
    call 155
    local.get 0
    i32.const 8
    i32.add
    call 82
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;190;) (type 15)
    (local i32 i64 i64 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 144
    i32.sub
    local.tee 0
    global.set 0
    call 31
    i32.const 1
    call 79
    call 77
    local.set 1
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              call 150
              call 86
              local.tee 2
              local.get 1
              i64.gt_u
              br_if 0 (;@5;)
              local.get 0
              local.get 1
              call 152
              local.get 0
              i32.load offset=88
              i32.eqz
              br_if 1 (;@4;)
              local.get 0
              local.get 1
              local.get 2
              call 146
              br_if 3 (;@2;)
              call 51
              local.set 3
              local.get 0
              i32.const 1
              i32.store offset=100
              local.get 0
              local.get 0
              i32.store offset=96
              loop  ;; label = @6
                local.get 0
                i32.const 104
                i32.add
                local.get 0
                i32.const 96
                i32.add
                call 44
                local.get 0
                i32.load8_u offset=136
                i32.const 2
                i32.eq
                br_if 3 (;@3;)
                local.get 3
                local.get 0
                i64.load offset=112
                call 118
                br 0 (;@6;)
              end
            end
            i64.const 0
            call 19
            br 3 (;@1;)
          end
          i64.const 1
          call 19
          br 2 (;@1;)
        end
        call 148
        call 86
        local.set 1
        local.get 0
        call 45
        i64.load
        local.set 2
        i32.const 0
        local.set 4
        i32.const 2
        i32.const 131988
        i32.const 0
        call 70
        local.tee 5
        call 126
        local.get 2
        local.get 1
        i64.add
        local.get 5
        call 125
        local.get 0
        local.get 3
        call 15
        local.tee 6
        i32.const 21
        i32.shl
        i32.const -16777216
        i32.and
        local.get 6
        i32.const 5
        i32.shl
        i32.const 16711680
        i32.and
        i32.or
        local.get 6
        i32.const 11
        i32.shr_u
        i32.const 65280
        i32.and
        local.get 6
        i32.const 27
        i32.shr_u
        i32.or
        i32.or
        i32.store
        local.get 5
        local.get 0
        i32.const 4
        call 13
        drop
        local.get 3
        call 15
        local.set 7
        block  ;; label = @3
          loop  ;; label = @4
            local.get 4
            i32.const 8
            i32.add
            local.tee 6
            local.get 7
            i32.gt_u
            br_if 1 (;@3;)
            local.get 0
            i64.const 0
            i64.store
            local.get 3
            local.get 4
            local.get 0
            i32.const 8
            call 113
            drop
            local.get 0
            i64.load
            local.tee 1
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
            i64.const 65280
            i64.and
            local.get 1
            i64.const 56
            i64.shr_u
            i64.or
            i64.or
            i64.or
            local.get 5
            call 125
            local.get 6
            local.set 4
            br 0 (;@4;)
          end
        end
        local.get 5
        call 20
        drop
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 2
        local.get 1
        i64.eq
        br_if 0 (;@2;)
        i64.const 3
        call 19
        br 1 (;@1;)
      end
      i64.const 4
      call 19
    end
    local.get 0
    i32.const 144
    i32.add
    global.set 0)
  (func (;191;) (type 15)
    call 31
    i32.const 0
    call 79
    call 150
    call 85)
  (func (;192;) (type 15)
    call 31
    i32.const 0
    call 79
    call 151
    call 85)
  (func (;193;) (type 15)
    (local i32 i32)
    call 31
    call 67
    i32.const 2
    call 79
    call 74
    local.set 0
    call 75
    local.set 1
    local.get 0
    call 168
    local.get 1
    call 59
    call 26
    drop)
  (func (;194;) (type 15)
    call 31
    i32.const 1
    call 79
    call 74
    call 168
    call 140
    call 38)
  (func (;195;) (type 15))
  (func (;196;) (type 15)
    call 173
    unreachable)
  (func (;197;) (type 2) (param i32 i32)
    call 196
    unreachable)
  (func (;198;) (type 2) (param i32 i32)
    call 196
    unreachable)
  (func (;199;) (type 8) (param i32 i32 i32) (result i32)
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
  (func (;200;) (type 8) (param i32 i32 i32) (result i32)
    (local i32 i32 i32)
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
        loop  ;; label = @3
          local.get 3
          local.get 1
          i32.store8
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
      local.tee 4
      i32.const -4
      i32.and
      local.tee 2
      i32.add
      local.set 3
      block  ;; label = @2
        local.get 2
        i32.const 1
        i32.lt_s
        br_if 0 (;@2;)
        local.get 1
        i32.const 255
        i32.and
        i32.const 16843009
        i32.mul
        local.set 2
        loop  ;; label = @3
          local.get 5
          local.get 2
          i32.store
          local.get 5
          i32.const 4
          i32.add
          local.tee 5
          local.get 3
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 4
      i32.const 3
      i32.and
      local.set 2
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
        i32.store8
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
  (func (;201;) (type 8) (param i32 i32 i32) (result i32)
    (local i32 i32 i32)
    i32.const 0
    local.set 3
    block  ;; label = @1
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      block  ;; label = @2
        loop  ;; label = @3
          local.get 0
          i32.load8_u
          local.tee 4
          local.get 1
          i32.load8_u
          local.tee 5
          i32.ne
          br_if 1 (;@2;)
          local.get 0
          i32.const 1
          i32.add
          local.set 0
          local.get 1
          i32.const 1
          i32.add
          local.set 1
          local.get 2
          i32.const -1
          i32.add
          local.tee 2
          i32.eqz
          br_if 2 (;@1;)
          br 0 (;@3;)
        end
      end
      local.get 4
      local.get 5
      i32.sub
      local.set 3
    end
    local.get 3)
  (func (;202;) (type 8) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call 201)
  (func (;203;) (type 8) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call 200)
  (func (;204;) (type 8) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call 199)
  (table (;0;) 1 1 funcref)
  (memory (;0;) 3)
  (global (;0;) (mut i32) (i32.const 131072))
  (global (;1;) i32 (i32.const 142116))
  (global (;2;) i32 (i32.const 142128))
  (export "memory" (memory 0))
  (export "init" (func 175))
  (export "batchTransferDcdtToken" (func 176))
  (export "moveRefundBatchToSafe" (func 177))
  (export "setWrappingContractAddress" (func 178))
  (export "setBridgeProxyContractAddress" (func 179))
  (export "addUnprocessedRefundTxToBatch" (func 180))
  (export "setDcdtSafeContractAddress" (func 181))
  (export "getWrappingContractAddress" (func 182))
  (export "getBridgeProxyContractAddress" (func 183))
  (export "getDcdtSafeContractAddress" (func 184))
  (export "setMaxTxBatchSize" (func 185))
  (export "setMaxTxBatchBlockDuration" (func 186))
  (export "getCurrentTxBatch" (func 187))
  (export "getFirstBatchAnyStatus" (func 188))
  (export "getBatch" (func 189))
  (export "getBatchStatus" (func 190))
  (export "getFirstBatchId" (func 191))
  (export "getLastBatchId" (func 192))
  (export "setMaxBridgedAmount" (func 193))
  (export "getMaxBridgedAmount" (func 194))
  (export "callBack" (func 195))
  (export "upgrade" (func 175))
  (export "__data_end" (global 1))
  (export "__heap_base" (global 2))
  (data (;0;) (i32.const 131072) "input too long\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00argument decode error (): too few argumentstoo many argumentswrong number of argumentsMultiDCDTNFTTransferDCDTNFTTransferDCDTTransfersync resultinput too shortManagedVec index out of range.itemstorage decode error (key: .lenbad array lengthMax tx batch size must be more than 0Max tx batch block duration must be more than 0new_max_tx_batch_sizevar argsgetTokensaddRefundBatchtransfersopt_new_addresschainSpecificToUniversalMappingInvalid token or amountInvalid unwrapping contract addressInvalid bridge proxy contract addressdepositwrapTokensinvalid valueinput out of rangelastBatchIdfirstBatchIdpendingBatchesmaxTxBatchSizemaxTxBatchBlockDurationunprocessedRefundTxstransferOverMaxAmounttransferPerformedEventwrappingContractAddressDcdtSafeContractAddressbridgeProxyContractAddresstransferFailedInvalidDestinationtransferFailedFrozenDestinationAccountmaxBridgedAmount\00Endpoint can only be called by ownerindex out of rangepanic occurredEmpty vec")
  (data (;1;) (i32.const 132068) "8\ff\ff\ff"))
