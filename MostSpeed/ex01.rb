prime = []
(1..100).each { |v|

  # 素数一覧へ初回は1が入力される
  if prime.empty?
    prime << v
  end

  add = v
  prime.each { |p|
    # 素数一覧の1はスキップ
    if p == 1
      next
    end

    # 素数2以上で割り切れる(=素数でないものがないかチェック)
    if v % p == 0
      add = nil
      break
    end
  }

  if not add.nil?
    prime << add
  end
  p prime
}
