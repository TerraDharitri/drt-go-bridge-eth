(module
  (type (;0;) (func (result i32)))
  (type (;1;) (func (param i32 i32) (result i32)))
  (type (;2;) (func (param i32 i32 i32) (result i32)))
  (type (;3;) (func (param i32)))
  (type (;4;) (func (param i32) (result i64)))
  (type (;5;) (func (param i32) (result i32)))
  (type (;6;) (func (param i32 i32)))
  (type (;7;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;8;) (func))
  (type (;9;) (func (result i64)))
  (type (;10;) (func (param i32 i32 i32)))
  (type (;11;) (func (param i32 i32 i32 i32)))
  (import "env" "mBufferNew" (func (;0;) (type 0)))
  (import "env" "mBufferAppend" (func (;1;) (type 1)))
  (import "env" "mBufferAppendBytes" (func (;2;) (type 2)))
  (import "env" "managedSignalError" (func (;3;) (type 3)))
  (import "env" "smallIntGetUnsignedArgument" (func (;4;) (type 4)))
  (import "env" "mBufferGetArgument" (func (;5;) (type 1)))
  (import "env" "mBufferGetLength" (func (;6;) (type 5)))
  (import "env" "getNumArguments" (func (;7;) (type 0)))
  (import "env" "signalError" (func (;8;) (type 6)))
  (import "env" "mBufferCopyByteSlice" (func (;9;) (type 7)))
  (import "env" "mBufferGetByteSlice" (func (;10;) (type 7)))
  (import "env" "mBufferSetBytes" (func (;11;) (type 2)))
  (import "env" "mBufferStorageLoad" (func (;12;) (type 1)))
  (import "env" "checkNoPayment" (func (;13;) (type 8)))
  (import "env" "managedGetMultiDCDTCallValue" (func (;14;) (type 3)))
  (import "env" "mBufferStorageStore" (func (;15;) (type 1)))
  (import "env" "mBufferFinish" (func (;16;) (type 5)))
  (func (;17;) (type 5) (param i32) (result i32)
    (local i32)
    call 0
    local.tee 1
    local.get 0
    call 1
    drop
    local.get 1)
  (func (;18;) (type 8)
    (local i32)
    i32.const 131113
    i32.const 23
    call 19
    local.tee 0
    i32.const 131072
    i32.const 7
    call 2
    drop
    local.get 0
    i32.const 131136
    i32.const 3
    call 2
    drop
    local.get 0
    i32.const 131229
    i32.const 16
    call 2
    drop
    local.get 0
    call 3
    unreachable)
  (func (;19;) (type 1) (param i32 i32) (result i32)
    (local i32)
    call 22
    local.tee 2
    local.get 0
    local.get 1
    call 11
    drop
    local.get 2)
  (func (;20;) (type 9) (result i64)
    i32.const 0
    call 4)
  (func (;21;) (type 0) (result i32)
    (local i32)
    i32.const 1
    call 22
    local.tee 0
    call 5
    drop
    block  ;; label = @1
      local.get 0
      call 6
      i32.const 32
      i32.eq
      br_if 0 (;@1;)
      call 18
      unreachable
    end
    local.get 0)
  (func (;22;) (type 0) (result i32)
    (local i32)
    i32.const 0
    i32.const 0
    i32.load offset=131296
    i32.const -1
    i32.add
    local.tee 0
    i32.store offset=131296
    local.get 0)
  (func (;23;) (type 3) (param i32)
    block  ;; label = @1
      call 7
      local.get 0
      i32.ne
      br_if 0 (;@1;)
      return
    end
    i32.const 131139
    i32.const 25
    call 8
    unreachable)
  (func (;24;) (type 2) (param i32 i32 i32) (result i32)
    (local i32 i32)
    local.get 0
    i32.load
    local.set 3
    call 0
    local.set 4
    block  ;; label = @1
      local.get 0
      i32.load offset=8
      local.get 3
      local.get 1
      local.get 4
      call 9
      br_if 0 (;@1;)
      local.get 0
      local.get 3
      local.get 1
      i32.add
      i32.store
      local.get 4
      return
    end
    local.get 2
    i32.const 131178
    i32.const 15
    call 25
    unreachable)
  (func (;25;) (type 10) (param i32 i32 i32)
    (local i32)
    i32.const 131198
    i32.const 27
    call 19
    local.tee 3
    local.get 0
    call 1
    drop
    local.get 3
    i32.const 131136
    i32.const 3
    call 2
    drop
    local.get 3
    local.get 1
    local.get 2
    call 2
    drop
    local.get 3
    call 3
    unreachable)
  (func (;26;) (type 7) (param i32 i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 3
    local.get 2
    call 10
    i32.const 0
    i32.ne)
  (func (;27;) (type 0) (result i32)
    (local i32)
    call 22
    local.tee 0
    i32.const 131264
    i32.const 0
    call 11
    drop
    local.get 0)
  (func (;28;) (type 5) (param i32) (result i32)
    (local i32)
    local.get 0
    call 22
    local.tee 1
    call 12
    drop
    local.get 1)
  (func (;29;) (type 5) (param i32) (result i32)
    (local i32 i32 i32 i64)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 1
    global.set 0
    local.get 1
    i64.const 0
    i64.store offset=8
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        call 28
        local.tee 2
        call 6
        local.tee 3
        i32.const 9
        i32.ge_u
        br_if 0 (;@2;)
        local.get 2
        i32.const 0
        local.get 1
        i32.const 8
        i32.add
        local.get 3
        i32.sub
        i32.const 8
        i32.add
        local.get 3
        call 26
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
        br_if 1 (;@1;)
        local.get 1
        i32.const 16
        i32.add
        global.set 0
        local.get 4
        i32.wrap_i64
        return
      end
      local.get 0
      i32.const 131164
      i32.const 14
      call 25
      unreachable
    end
    local.get 0
    i32.const 131164
    i32.const 14
    call 25
    unreachable)
  (func (;30;) (type 1) (param i32 i32) (result i32)
    (local i32)
    global.get 0
    i32.const 16
    i32.sub
    local.tee 2
    global.set 0
    local.get 0
    call 17
    local.tee 0
    i32.const 131193
    i32.const 5
    call 2
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
    call 2
    drop
    local.get 2
    i32.const 16
    i32.add
    global.set 0
    local.get 0)
  (func (;31;) (type 3) (param i32)
    local.get 0
    i32.const 131178
    i32.const 15
    call 25
    unreachable)
  (func (;32;) (type 11) (param i32 i32 i32 i32)
    (local i32 i32 i32)
    local.get 0
    i32.load
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 0
                  i32.load8_u offset=16
                  br_if 0 (;@7;)
                  local.get 0
                  i32.load offset=8
                  local.tee 5
                  call 6
                  local.tee 6
                  i32.const 10000
                  i32.gt_u
                  br_if 1 (;@6;)
                  i32.const 0
                  i32.load8_u offset=141304
                  i32.const 255
                  i32.and
                  br_if 1 (;@6;)
                  i32.const 0
                  local.get 6
                  i32.store offset=141300
                  i32.const 0
                  i32.const 1
                  i32.store8 offset=141304
                  local.get 5
                  i32.const 0
                  i32.const 131300
                  local.get 6
                  call 26
                  drop
                  local.get 0
                  i32.const 1
                  i32.store8 offset=16
                end
                local.get 4
                local.get 2
                i32.add
                local.tee 5
                i32.const 0
                i32.load offset=141300
                i32.le_u
                br_if 1 (;@5;)
                br 5 (;@1;)
              end
              local.get 0
              i32.const 0
              i32.store8 offset=16
              local.get 5
              local.get 4
              local.get 1
              local.get 2
              call 26
              br_if 4 (;@1;)
              local.get 4
              local.get 2
              i32.add
              local.set 5
              br 1 (;@4;)
            end
            local.get 5
            local.get 4
            i32.lt_u
            br_if 1 (;@3;)
            local.get 5
            i32.const 10000
            i32.gt_u
            br_if 2 (;@2;)
            local.get 1
            local.get 4
            i32.const 131300
            i32.add
            local.get 2
            call 48
            drop
          end
          local.get 0
          local.get 5
          i32.store
          return
        end
        local.get 4
        local.get 5
        call 33
        unreachable
      end
      local.get 5
      call 34
      unreachable
    end
    local.get 3
    call 31
    unreachable)
  (func (;33;) (type 6) (param i32 i32)
    local.get 0
    local.get 1
    call 46
    unreachable)
  (func (;34;) (type 3) (param i32)
    local.get 0
    call 45
    unreachable)
  (func (;35;) (type 3) (param i32)
    (local i32 i32)
    i32.const 131245
    i32.const 16
    call 19
    local.tee 1
    call 17
    local.tee 2
    i32.const 131225
    i32.const 4
    call 2
    drop
    local.get 0
    local.get 2
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func (;36;) (type 8)
    call 37
    unreachable)
  (func (;37;) (type 8)
    i32.const 131282
    i32.const 14
    call 8
    unreachable)
  (func (;38;) (type 8)
    i32.const 0
    call 23)
  (func (;39;) (type 8)
    call 13
    i32.const 0
    call 23)
  (func (;40;) (type 8)
    (local i32 i64 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 0
    global.set 0
    i32.const 2
    call 23
    call 20
    local.set 1
    call 21
    local.set 2
    i32.const -21
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        i32.const 0
        i32.load8_u offset=141308
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
      i32.store8 offset=141308
      i32.const -21
      call 14
    end
    block  ;; label = @1
      block  ;; label = @2
        local.get 3
        call 6
        i32.const -16
        i32.and
        i32.const 16
        i32.ne
        br_if 0 (;@2;)
        i32.const 0
        local.set 4
        local.get 3
        call 6
        local.set 5
        local.get 0
        i32.const 16
        i32.add
        i32.const 8
        i32.add
        local.set 6
        i32.const 1
        local.set 7
        i32.const 0
        local.set 8
        loop  ;; label = @3
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
          local.get 0
          i64.const 0
          i64.store offset=16
          local.get 3
          local.get 4
          local.get 0
          i32.const 16
          i32.add
          i32.const 16
          call 26
          drop
          block  ;; label = @4
            local.get 7
            i32.const 1
            i32.and
            i32.eqz
            br_if 0 (;@4;)
            local.get 0
            i32.load offset=16
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
            local.set 10
            i32.const 1
            local.set 8
            i32.const 0
            local.set 7
            local.get 9
            local.set 4
            br 1 (;@3;)
          end
        end
        local.get 8
        call 41
        unreachable
      end
      i32.const 131079
      i32.const 34
      call 8
      unreachable
    end
    local.get 0
    i32.const 8
    i32.add
    call 35
    local.get 0
    i32.load offset=8
    local.get 0
    i32.load offset=12
    local.tee 3
    call 29
    i32.const 1
    i32.add
    local.tee 7
    call 30
    local.set 5
    i32.const 131264
    i32.const 0
    call 19
    local.set 4
    local.get 0
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
    i64.store offset=16
    local.get 4
    local.get 0
    i32.const 16
    i32.add
    i32.const 8
    call 2
    drop
    local.get 4
    local.get 2
    call 1
    drop
    local.get 0
    local.get 10
    call 6
    local.tee 9
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
    i32.store offset=16
    local.get 4
    local.get 0
    i32.const 16
    i32.add
    i32.const 4
    call 2
    drop
    local.get 4
    local.get 10
    call 1
    drop
    local.get 5
    local.get 4
    call 15
    drop
    local.get 0
    local.get 7
    i32.store8 offset=23
    local.get 0
    i32.const 0
    i32.store offset=16
    local.get 0
    local.get 7
    i64.extend_i32_u
    local.tee 1
    i64.const 24
    i64.shr_u
    i32.wrap_i64
    local.tee 4
    i32.store8 offset=20
    local.get 0
    local.get 1
    i64.const 16
    i64.shr_u
    i32.wrap_i64
    local.tee 9
    i32.store8 offset=21
    local.get 0
    local.get 1
    i64.const 8
    i64.shr_u
    i32.wrap_i64
    local.tee 10
    i32.store8 offset=22
    local.get 3
    local.get 0
    i32.const 16
    i32.add
    i32.const 4
    i32.const 5
    local.get 4
    select
    local.get 9
    local.get 4
    i32.or
    i32.const 255
    i32.and
    i32.eqz
    local.tee 4
    i32.add
    i32.const 0
    local.get 4
    local.get 10
    i32.const 255
    i32.and
    select
    local.tee 4
    i32.add
    i32.const 0
    local.get 4
    local.get 7
    select
    i32.add
    local.tee 4
    i32.add
    i32.const 8
    local.get 4
    i32.sub
    call 19
    call 15
    drop
    local.get 0
    i32.const 32
    i32.add
    global.set 0)
  (func (;41;) (type 3) (param i32)
    call 44
    unreachable)
  (func (;42;) (type 8)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i64 i32)
    global.get 0
    i32.const 64
    i32.sub
    local.tee 0
    global.set 0
    call 13
    i32.const 0
    call 23
    call 27
    local.set 1
    local.get 0
    i32.const 24
    i32.add
    call 35
    local.get 0
    i32.load offset=28
    call 29
    local.set 2
    i32.const 1
    local.set 3
    i32.const 0
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          loop  ;; label = @4
            local.get 3
            local.set 5
            local.get 4
            i32.const 1
            i32.and
            br_if 1 (;@3;)
            local.get 5
            local.get 2
            i32.gt_u
            br_if 1 (;@3;)
            local.get 0
            i32.const 16
            i32.add
            call 35
            local.get 0
            i32.load offset=16
            local.set 6
            local.get 0
            i32.load offset=20
            call 29
            local.get 5
            i32.lt_u
            br_if 2 (;@2;)
            local.get 5
            local.get 2
            i32.ge_u
            local.set 4
            local.get 5
            local.get 5
            local.get 2
            i32.lt_u
            i32.add
            local.set 3
            local.get 6
            local.get 5
            call 30
            i32.const -25
            call 12
            drop
            i32.const -25
            call 6
            i32.eqz
            br_if 0 (;@4;)
            local.get 0
            i32.const 8
            i32.add
            call 35
            local.get 0
            i32.load offset=8
            local.get 5
            call 30
            local.tee 5
            call 28
            local.tee 7
            call 6
            local.set 6
            local.get 0
            i32.const 0
            i32.store8 offset=48
            local.get 0
            local.get 6
            i32.store offset=44
            local.get 0
            local.get 7
            i32.store offset=40
            local.get 0
            local.get 6
            i32.store offset=36
            local.get 0
            i32.const 0
            i32.store offset=32
            local.get 0
            i64.const 0
            i64.store offset=56
            local.get 0
            i32.const 32
            i32.add
            local.get 0
            i32.const 56
            i32.add
            i32.const 8
            local.get 5
            call 32
            local.get 0
            i64.load offset=56
            local.set 8
            local.get 0
            i32.const 32
            i32.add
            i32.const 32
            local.get 5
            call 24
            local.set 9
            local.get 0
            i32.const 0
            i32.store offset=56
            local.get 0
            i32.const 32
            i32.add
            local.get 0
            i32.const 56
            i32.add
            i32.const 4
            local.get 5
            call 32
            local.get 0
            i32.const 32
            i32.add
            local.get 0
            i32.load offset=56
            local.tee 6
            i32.const 24
            i32.shl
            local.get 6
            i32.const 65280
            i32.and
            i32.const 8
            i32.shl
            i32.or
            local.get 6
            i32.const 8
            i32.shr_u
            i32.const 65280
            i32.and
            local.get 6
            i32.const 24
            i32.shr_u
            i32.or
            i32.or
            local.get 5
            call 24
            local.set 7
            local.get 0
            i32.load offset=36
            local.get 0
            i32.load offset=32
            i32.ne
            br_if 3 (;@1;)
            block  ;; label = @5
              local.get 0
              i32.load8_u offset=48
              i32.eqz
              br_if 0 (;@5;)
              i32.const 0
              i32.const 0
              i32.store offset=141300
              i32.const 0
              i32.const 0
              i32.store8 offset=141304
            end
            call 27
            call 17
            local.set 5
            local.get 0
            local.get 8
            i64.store offset=32
            local.get 5
            local.get 0
            i32.const 32
            i32.add
            i32.const 8
            call 2
            drop
            local.get 5
            local.get 9
            call 1
            drop
            local.get 0
            local.get 7
            call 6
            local.tee 6
            i32.const 24
            i32.shl
            local.get 6
            i32.const 65280
            i32.and
            i32.const 8
            i32.shl
            i32.or
            local.get 6
            i32.const 8
            i32.shr_u
            i32.const 65280
            i32.and
            local.get 6
            i32.const 24
            i32.shr_u
            i32.or
            i32.or
            i32.store offset=32
            local.get 5
            local.get 0
            i32.const 32
            i32.add
            i32.const 4
            call 2
            drop
            local.get 5
            local.get 7
            call 1
            drop
            local.get 0
            local.get 5
            i32.const 24
            i32.shl
            local.get 5
            i32.const 65280
            i32.and
            i32.const 8
            i32.shl
            i32.or
            local.get 5
            i32.const 8
            i32.shr_u
            i32.const 65280
            i32.and
            local.get 5
            i32.const 24
            i32.shr_u
            i32.or
            i32.or
            i32.store offset=32
            local.get 1
            local.get 0
            i32.const 32
            i32.add
            i32.const 4
            call 2
            drop
            br 0 (;@4;)
          end
        end
        local.get 1
        call 6
        local.set 6
        i32.const 0
        local.set 5
        block  ;; label = @3
          loop  ;; label = @4
            local.get 5
            i32.const 4
            i32.add
            local.tee 2
            local.get 6
            i32.gt_u
            br_if 1 (;@3;)
            local.get 0
            i32.const 0
            i32.store offset=32
            local.get 1
            local.get 5
            local.get 0
            i32.const 32
            i32.add
            i32.const 4
            call 26
            drop
            local.get 0
            i32.load offset=32
            local.tee 5
            i32.const 24
            i32.shl
            local.get 5
            i32.const 65280
            i32.and
            i32.const 8
            i32.shl
            i32.or
            local.get 5
            i32.const 8
            i32.shr_u
            i32.const 65280
            i32.and
            local.get 5
            i32.const 24
            i32.shr_u
            i32.or
            i32.or
            call 16
            drop
            local.get 2
            local.set 5
            br 0 (;@4;)
          end
        end
        local.get 0
        i32.const 64
        i32.add
        global.set 0
        return
      end
      i32.const 131264
      i32.const 18
      call 8
      unreachable
    end
    local.get 5
    i32.const 131164
    i32.const 14
    call 25
    unreachable)
  (func (;43;) (type 8))
  (func (;44;) (type 8)
    call 36
    unreachable)
  (func (;45;) (type 3) (param i32)
    call 44
    unreachable)
  (func (;46;) (type 6) (param i32 i32)
    call 44
    unreachable)
  (func (;47;) (type 2) (param i32 i32 i32) (result i32)
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
  (func (;48;) (type 2) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call 47)
  (table (;0;) 1 1 funcref)
  (memory (;0;) 3)
  (global (;0;) (mut i32) (i32.const 131072))
  (global (;1;) i32 (i32.const 141309))
  (global (;2;) i32 (i32.const 141312))
  (export "memory" (memory 0))
  (export "callPayable" (func 38))
  (export "callNonPayable" (func 39))
  (export "callPayableWithParams" (func 40))
  (export "getCalledDataParams" (func 42))
  (export "callBack" (func 43))
  (export "upgrade" (func 39))
  (export "init" (func 39))
  (export "__data_end" (global 1))
  (export "__heap_base" (global 2))
  (data (;0;) (i32.const 131072) "addressincorrect number of DCDT transfersargument decode error (): wrong number of argumentsinput too longinput too short.itemstorage decode error (key: .lenbad array lengthcalledDataParams\00\00\00index out of rangepanic occurred")
  (data (;1;) (i32.const 131296) "8\ff\ff\ff"))
